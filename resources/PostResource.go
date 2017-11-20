package resources

import (
	"github.com/emicklei/go-restful"
	"../dao"
	"strconv"
	"../utils"
	"../domain"
	"os"
	"github.com/franela/goreq"
)

type PostResource struct {
	Dao dao.PostDao

}

func (postRes PostResource) Register(container* restful.Container) {
	ws := new(restful.WebService)
	ws.
		Path("/posts").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("/{post-id}").To(postRes.FindPost))
	ws.Route(ws.POST("").Filter(postRes.CheckCaptcha).To(postRes.CreatePost).Reads(domain.Post{}))
	ws.Route(ws.GET("").To(postRes.GetPosts))

	container.Add(ws)
}

func (postRes PostResource) FindPost(request *restful.Request, response *restful.Response) {
	postId, parseError := strconv.Atoi(request.PathParameter("post-id"))
	if parseError != nil {
		response.WriteHeader(400)
		response.WriteEntity(utils.ApiResponse{Message: "Error"})
		return
	}
	post, err := postRes.Dao.FindOne(uint(postId))
	apiResponse := utils.ApiResponse{Message: "Ok", Data: &post}
	if err == nil {
		response.WriteEntity(apiResponse)
	} else {
		response.WriteHeader(404)
		response.WriteEntity(utils.ApiResponse{Message: "Not found"})
	}
	
}

func (postRes PostResource) CheckCaptcha(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	token := os.Getenv("RECAPTCHA_TOKEN")
	captchaResponse, noHeaderError := req.BodyParameter("")
	if noHeaderError != nil {
		apiResponse := utils.ApiResponse{Message: "Error"}
		resp.WriteHeader(400)
		resp.WriteEntity(apiResponse)
		return
	}

	res, err := goreq.Request{
		Method: "POST",
		Uri: "https://www.google.com/recaptcha/api/siteverify",
		Body: "{\"secret\":\"" + token + "\",\"response\":\"" + captchaResponse + "\"}",
	}.Do()

	if err != nil {
		println(res.Body)
	}
}

func (postRes PostResource) CreatePost(request *restful.Request, response *restful.Response) {
	var post domain.Post
	request.ReadEntity(&post)
	if len(post.Content) == 0 {
		response.WriteEntity(utils.ApiResponse{Message: "Error", Data: &post})
		return
	}
	post = postRes.Dao.Create(post.Content)
	apiResponse := utils.ApiResponse{Message: "Ok", Data: &post}
	response.WriteEntity(apiResponse)

}

func (postRes PostResource) GetPosts(request *restful.Request, response *restful.Response) {
	offset, parseError1 := strconv.Atoi(request.QueryParameter("offset"))
	limit, parseError2 := strconv.Atoi(request.QueryParameter("limit"))
	if parseError1 != nil || parseError2 != nil {
			println(offset, limit)

		response.WriteHeader(400)
		response.WriteEntity(utils.ApiResponse{Message: "Error"})
		return
	}
	posts, err := postRes.Dao.FindNAmount(int(offset), int(limit))
	apiResponse := utils.ApiResponseArray{Message: "Ok", Data: posts}
	if err == nil {
		response.WriteEntity(apiResponse)
	} else {
		response.WriteHeader(404)
		response.WriteEntity(utils.ApiResponse{Message: "Not found"})
	}
	
}