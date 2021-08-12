func TestGetByID(t *testing.T) {
	var mockArticle models.Article 
	err := faker.FakeData(&mockArticle) 
	assert.NoError(t, err) 
	mockUCase := new(mocks.ArticleUsecase) 
	num := int(mockArticle.ID) 
	mockUCase.On(“GetByID”, int64(num)).Return(&mockArticle, nil) 
	e := echo.New() 
	req, err := http.NewRequest(echo.GET, “/article/” +  
				strconv.Itoa(int(num)), strings.NewReader(“”)) 
	assert.NoError(t, err) 
	rec := httptest.NewRecorder() 
	c := e.NewContext(req, rec) 
	c.SetPath(“article/:id”) 
	c.SetParamNames(“id”) 
	c.SetParamValues(strconv.Itoa(num)) 
	handler:= articleHttp.ArticleHandler{
			   AUsecase: mockUCase,
			   Helper: httpHelper.HttpHelper{}
	} 
	handler.GetByID(c) 
	assert.Equal(t, http.StatusOK, rec.Code) 
	mockUCase.AssertCalled(t, “GetByID”, int64(num))
   }