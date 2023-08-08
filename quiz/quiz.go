package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"net/http"
	"quiz_assessment/data"
)

var getQuestionsCmd = &cobra.Command{
	Use:   "get-questions",
	Short: "Get quiz questions from the API",
	Run:   getQuestions,
}
var startQuizCmd = &cobra.Command{
	Use:   "start-quiz",
	Short: "Start quiz questions from the API",
	Run:   startQuiz,
}

func init() {
	rootCmd.AddCommand(getQuestionsCmd)
	rootCmd.AddCommand(startQuizCmd)
}

func getQuestions(cmd *cobra.Command, args []string) {
	resp, err := http.Get("http://localhost:8080/questions")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	var questions []data.Question
	err = json.NewDecoder(resp.Body).Decode(&questions)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, question := range questions {
		fmt.Println(question.Text)
		for i, option := range question.Options {
			fmt.Printf("%d. %s\n", i+1, option)
		}
	}
}

func startQuiz(cmd *cobra.Command, args []string) {
	resp, err := http.Get("http://localhost:8080/questions")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	var questions []data.Question
	err = json.NewDecoder(resp.Body).Decode(&questions)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	correctAnswers := 0
	for _, question := range questions {
		fmt.Println(question.Text)
		for i, option := range question.Options {
			fmt.Printf("%d. %s\n", i+1, option)
		}

		var userChoice int
		fmt.Print("Your answer (1-4): ")
		_, err := fmt.Scan(&userChoice)
		if err != nil || userChoice < 1 || userChoice > 4 {
			fmt.Println("Invalid choice. Skipping to the next question.")
			continue
		}

		if userChoice-1 == question.CorrectIdx {
			correctAnswers++
		}
	}

	totalQuestions := len(questions)
	fmt.Printf("You answered %d out of %d questions correctly.\n", correctAnswers, totalQuestions)

	scorePercentage := (correctAnswers * 100) / totalQuestions

	calculateAndCompareScores(scorePercentage)
	submitScore(scorePercentage)
}
func submitScore(score int) {
	reqBody, _ := json.Marshal(score)
	resp, err := http.Post("http://localhost:8080/submit", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		fmt.Println("Error submitting score:", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		var quizResults []int
		err := json.NewDecoder(resp.Body).Decode(&quizResults)
		if err != nil {
			fmt.Println("Error decoding response:", err)
			return
		}

	} else {
		fmt.Println("Score submission failed with status:", resp.Status)
	}

}

func calculateAndCompareScores(userScore int) {
	resp, err := http.Get("http://localhost:8080/results")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		var quizResults []int
		err := json.NewDecoder(resp.Body).Decode(&quizResults)
		if err != nil {
			fmt.Println("Error decoding response:", err)
			return
		}

		// Compare userScore with quizResults and provide feedback
		numBetter := 0
		totalScores := len(quizResults)

		for _, score := range quizResults {
			if userScore > score {
				numBetter++
			}
		}

		if len(quizResults) == 0 {
			fmt.Printf("You are first quizzer")
		} else {
			percentageBetter := (numBetter * 100) / totalScores
			fmt.Printf("You were better than %d%% of all quizzers.\n", percentageBetter)
		}

	} else {
		fmt.Println("Score submission failed with status:", resp.Status)
	}

}

var rootCmd = &cobra.Command{
	Use:   "quiz",
	Short: "A simple quiz CLI",
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
