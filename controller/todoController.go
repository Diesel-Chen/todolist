package controller

import (
	"log"
	"net/http"
	"todolist/model"

	"github.com/gin-gonic/gin"
)

type PostMessage struct {
	Title     string `json:"title" form:"title"`
	Completed int    `json:"complete" form:"complete"`
}

type PutMessage struct {
	Title     string `json:"title" form:"title"`
	Completed int    `json:"complete" form:"complete"`
}

// AddTodo 是添加一条todo项目
func AddTodo(c *gin.Context) {

	var postMessage PostMessage

	c.BindJSON(&postMessage)

	//from 表单中的json
	log.Println(postMessage.Title, postMessage.Completed)
	todo := model.TodoModel{Title: postMessage.Title, Completed: postMessage.Completed}

	model.DBEngin.Save(&todo)
	c.JSON(http.StatusCreated, gin.H{
		"status":     http.StatusCreated,
		"message":    "成功地添加了一个代办项目",
		"resourceId": todo.ID,
	})

}

// updateTitleTodo 是修改todo项目名
func UpdateTitleTodo(c *gin.Context) {

	var putMessage PutMessage

	c.BindJSON(&putMessage)

	var todo model.TodoModel

	todoID := c.Param("id")

	model.DBEngin.First(&todo, todoID)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "没有 todo 项目",
		})
		return
	}

	model.DBEngin.Model(&todo).Update("title", putMessage.Title)
	// completed, _ := strconv.Atoi(c.PostForm("completed"))
	// model.DBEngin.Model(&todo).Update("completed", completed)
	model.DBEngin.Model(&todo).Update("completed", 0)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "修改todo名称成功!",
	})

}

// SelectAllTodo 查询全部的todo项目
func SelectAllTodo(c *gin.Context) {
	var todos []model.TodoModel
	var _todos []model.TransformedTodo

	model.DBEngin.Find(&todos)

	if len(todos) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "没有 todo 项目",
		})
		return
	}

	//对返回给客户端的数据进行处理
	for _, item := range todos {
		completed := false
		if item.Completed == 1 {
			completed = true
		} else {
			completed = false
		}
		_todos = append(_todos, model.TransformedTodo{ID: item.ID, Title: item.Title, Completed: completed})
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   _todos,
	})
}

// SelectSingleTodo 根据id查询具体的某项 todo
func SelectSingleTodo(c *gin.Context) {
	var todo model.TodoModel

	todoID := c.Param("id")

	model.DBEngin.First(&todo, todoID)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "没有 todo 项目",
		})
		return
	}

	completed := false
	if todo.Completed == 1 {
		completed = true
	} else {
		completed = false
	}

	_todo := model.TransformedTodo{ID: todo.ID, Title: todo.Title, Completed: completed}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   _todo,
	})
}

//SelectCompleted 查询已经完成的 todolist
func SelectCompleted(c *gin.Context) {
	var todos []model.TodoModel
	var _todos []model.TransformedTodo

	model.DBEngin.Where("completed=?", "1").Find(&todos)

	if len(todos) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "没有 todo 项目",
		})
		return
	}

	//对返回给客户端的数据进行处理
	for _, item := range todos {
		completed := false
		if item.Completed == 1 {
			completed = true
		} else {
			completed = false
		}
		_todos = append(_todos, model.TransformedTodo{ID: item.ID, Title: item.Title, Completed: completed})
	}
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   _todos,
	})
}

// UpdateStateTodo  //根据某项 todo id修改 项目已完成。
func UpdateStateTodo(c *gin.Context) {
	var todo model.TodoModel

	todoID := c.Param("id")

	model.DBEngin.First(&todo, todoID)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "没有 todo 项目",
		})
		return
	}

	// model.DBEngin.Model(&todo).Update("title", c.PostForm("title"))
	// completed, _ := strconv.Atoi(c.PostForm("completed"))
	// model.DBEngin.Model(&todo).Update("completed", completed)
	model.DBEngin.Model(&todo).Update("completed", 1)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "修改成功!",
	})
}

// DeleteTodo 根据id删除todo 项目
func DeleteTodo(c *gin.Context) {
	var todo model.TodoModel
	todoID := c.Param("id")

	model.DBEngin.First(&todo, todoID)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "没有找到 todo 项目!"})
		return
	}

	model.DBEngin.Unscoped().Delete(&todo)
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "删除成功!",
	})
}
