package quiz

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"techtestify/internal/db"
	"techtestify/internal/models"
)

type CreateTestInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

type CreateQuestionInput struct {
	Text    string `json:"text" binding:"required"`
	OptionA string `json:"optionA" binding:"required"`
	OptionB string `json:"optionB" binding:"required"`
	OptionC string `json:"optionC" binding:"required"`
	OptionD string `json:"optionD" binding:"required"`
	Answer  string `json:"answer" binding:"required,oneof=A B C D"`
}

type SubmitAnswersInput struct {
	Answers map[string]string `json:"answers" binding:"required"`
}

func CreateTest(c *gin.Context) {
	var input CreateTestInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	userID := c.GetUint("user_id")
	test := models.Test{
		Title:       input.Title,
		Description: input.Description,
		CreatedBy:   userID,
	}

	if err := db.DB.Create(&test).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create test"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Test created", "test": test})
}

func AddQuestion(c *gin.Context) {
	testIDStr := c.Param("id")
	testID, err := strconv.Atoi(testIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid test ID"})
		return
	}

	var input CreateQuestionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	question := models.Question{
		TestID:  uint(testID),
		Text:    input.Text,
		OptionA: input.OptionA,
		OptionB: input.OptionB,
		OptionC: input.OptionC,
		OptionD: input.OptionD,
		Answer:  input.Answer,
	}

	if err := db.DB.Create(&question).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not add question"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Question added", "question": question})
}

func GetAllTests(c *gin.Context) {
	var tests []models.Test
	if err := db.DB.Preload("Questions").Find(&tests).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch tests"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tests": tests})
}

func SubmitTest(c *gin.Context) {
	testIDStr := c.Param("id")
	testID, err := strconv.Atoi(testIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid test ID"})
		return
	}

	var input SubmitAnswersInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var questions []models.Question
	if err := db.DB.Where("test_id = ?", testID).Find(&questions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch questions"})
		return
	}

	total := len(questions)
	score := 0
	correctAnswers := make(map[string]string)

	for _, q := range questions {
		correctAnswers[strconv.Itoa(int(q.ID))] = q.Answer
		if ans, ok := input.Answers[strconv.Itoa(int(q.ID))]; ok && ans == q.Answer {
			score++
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"score":           score,
		"total":           total,
		"correct_answers": correctAnswers,
	})
}
