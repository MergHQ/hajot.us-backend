package resources

import (
	"github.com/emicklei/go-restful"
	"io"
	"../dao"
	"../domain"
)

type PostResource struct {
	dao dao.PostDao

}

type Post struct {
	Id int
	content string
	timestamp string
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

}

func (postRes PostResource) CreatePost(request *restful.Request, response *restful.Response) {

}

func (postRes PostResource) GetPosts(request *restful.Request, response *restful.Response) {

}