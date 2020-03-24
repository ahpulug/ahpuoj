package router

import (
	"ahpuoj/controller"

	"github.com/gin-gonic/gin"
)

func ApiAdminRouter(g *gin.RouterGroup) {
	// 竞赛
	g.GET("/contest/:id/problem/:problemid/solutions", controller.GetContestProblemSolutions)
	g.POST("/contest", controller.StoreContest)
	g.GET("/contests", controller.IndexContest)
	g.GET("/allcontests", controller.GetAllContests)
	g.GET("/contest/:id", controller.ShowContest)
	g.PUT("/contest/:id", controller.UpdateContest)
	g.DELETE("/contest/:id", controller.DeleteContest)
	g.PUT("/contest/:id/status", controller.ToggleContestStatus)
	g.GET("/contest/:id/users", controller.IndexContestUser)
	g.DELETE("/contest/:id/user/:userid", controller.DeleteContestUser)
	g.GET("/contest/:id/teams", controller.IndexContestTeamWithUser)
	g.POST("/contest/:id/team/:teamid", controller.AddContestTeam)
	g.DELETE("/contest/:id/team/:teamid", controller.DeleteContestTeam)
	g.POST("/contest/:id/team/:teamid/users", controller.AddContestTeamUsers)
	g.POST("/contest/:id/team/:teamid/allusers", controller.AddContestTeamAllUsers)
	g.DELETE("/contest/:id/team/:teamid/user/:userid", controller.DeleteContestTeamUser)
	// 生成器
	g.POST("/generator/compete", controller.CompeteAccountGenerator)
	g.POST("/generator/user", controller.UserAccountGenerator)
	// 图片
	g.POST("/image", controller.StoreImage)
	// 讨论
	g.PUT("/issue/:id/status", controller.ToggleIssueStatus)
	g.PUT("/reply/:id/status", controller.ToggleReplyStatus)
	// 新闻
	g.GET("/news", controller.IndexNew)
	g.POST("/new", controller.StoreNew)
	g.GET("/new/:id", controller.ShowNew)
	g.PUT("/new/:id", controller.UpdateNew)
	g.DELETE("/new/:id", controller.DeleteNew)
	g.PUT("/new/:id/status", controller.ToggleNewStatus)
	g.PUT("/new/:id/topstatus", controller.ToggleNewTopStatus)
	// 问题
	g.POST("/problem", controller.StoreProblem)
	g.GET("/problems", controller.IndexProblem)
	g.GET("/problem/:id", controller.ShowProblem)
	g.PUT("/problem/:id", controller.UpdateProblem)
	g.DELETE("/problem/:id", controller.DeleteProblem)
	g.PUT("/problem/:id/status", controller.ToggleProblemStatus)
	g.GET("/problem/:id/datas", controller.IndexProblemData)
	g.GET("/problem/:id/data/:filename", controller.GetProblemData)
	g.GET("/download/problem/:id/data/:filename", controller.DownloadProblemData)
	g.POST("/problem/:id/data", controller.AddProblemData)
	g.POST("/problem/:id/datafile", controller.AddProblemDataFile)
	g.PUT("/problem/:id/data/:filename", controller.EditProblemData)
	g.DELETE("/problem/:id/data/:filename", controller.DeleteProblemData)
	g.PUT("/solution/:id/judgestatus", controller.RejudgeSolution)
	g.PUT("/problem/:id/judgestatus", controller.RejudgeProblem)
	g.PUT("/problem/:id/movement/:newid", controller.ReassignProblem)
	// 问题集
	g.POST("/problemset", controller.ImportProblemSet)
	// 系列赛事
	g.POST("/series", controller.StoreSeries)
	g.GET("/serieses", controller.IndexSeries)
	g.GET("/series/:id", controller.ShowSeries)
	g.GET("/series/:id/contests", controller.IndexSeriesContest)
	g.PUT("/series/:id", controller.UpdateSeries)
	g.DELETE("/series/:id", controller.DeleteSeries)
	g.PUT("/series/:id/status", controller.ToggleSeriesStatus)
	g.POST("/series/:id/contest/:contestid", controller.AddSeriesContest)
	g.DELETE("/series/:id/contest/:contestid", controller.DeleteSeriesContest)
	// 设置
	g.GET("/settings", controller.GetSettings)
	g.PUT("/settings", controller.SetSettings)
	// 标签
	g.POST("/tag", controller.StoreTag)
	g.GET("/tags", controller.IndexTag)
	g.GET("/alltags", controller.GetAllTags)
	g.PUT("/tag/:id", controller.UpdateTag)
	g.DELETE("/tag/:id", controller.DeleteTag)
	// 用户
	g.GET("/users", controller.IndexUser)
	g.PUT("/user/:id/status", controller.ToggleUserStatus)
	g.PUT("/user/:id/pass", controller.ChangeUserPass)
	// 首页
	g.GET("/submitstatistic", controller.GetSubmitStatistic)
	// 团队
	g.POST("/team", controller.StoreTeam)
	g.POST("/team/:id/users", controller.AddTeamUsers)
	g.GET("/teams", controller.IndexTeam)
	g.GET("/allteams", controller.GetAllTeams) // 在竞赛作业人员配置中获取全部团队的接口
	g.GET("/team/:id", controller.GetTeam)
	g.GET("/team/:id/users", controller.IndexTeamUser)
	g.PUT("/team/:id", controller.UpdateTeam)
	g.DELETE("/team/:id/user/:userid", controller.DeleteTeamUser)
	g.DELETE("/team/:id", controller.DeleteTeam)
}
