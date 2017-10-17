package resources

import (
	"github.com/emicklei/go-restful"
	"../dao"
	"strconv"
	"../utils"
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
	ws.Route(ws.POST("").To(postRes.CreatePost))
	ws.Route(ws.GET("").To(postRes.GetPosts))

	container.Add(ws)
}

func (postRes PostResource) FindPost(request *restful.Request, response *restful.Response) {
	postId, parseError := strconv.ParseUint(request.PathParameter("post-id"), 10, 1)
	if parseError != nil {
		panic("string parse error")
	}
	user, err := postRes.Dao.FindOne(uint(postId))
	apiResponse := utils.ApiResponse{Message: "Ok", Data: user}
	if err == nil {
		response.WriteEntity(apiResponse)
	}
	
}

func (postRes PostResource) CreatePost(request *restful.Request, response *restful.Response) {

}

func (postRes PostResource) GetPosts(request *restful.Request, response *restful.Response) {

}