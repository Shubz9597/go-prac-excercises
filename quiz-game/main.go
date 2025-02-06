package main

import (
	"csvReader/csvRead"
	"flag"
	"fmt"
	"math/rand"
	"time"
)

type ProgramParams struct {
	file     string
	quizTime int
	shuffle  bool
}

func shuffleQuestions(ques [][]string) (shuffleQues [][]string) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	shuffleQues = make([][]string, len(ques))

	for i, index := range r.Perm(len(ques)) {
		shuffleQues[i] = ques[index]
	}
	return
}

func quizGame(questions [][]string, timer *time.Timer) (int, int) {

	correct_answers := 0

	for index, value := range questions {
		fmt.Printf("Question Number %d: %s = ", index+1, value[0])
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println("\nTime's Up!!!!")
			return correct_answers, len(questions)
		case answer := <-answerCh:
			if answer == value[1] {
				correct_answers++
			}
		}
	}

	return correct_answers, len(questions)
}

func main() {
	var params ProgramParams

	flag.StringVar(&params.file, "file", "problems", "File Name used to give the quiz")
	flag.IntVar(&params.quizTime, "t", 30, "Time in seconds to solve the given quiz")
	flag.BoolVar(&params.shuffle, "shuffle", false, "Shuffle the Questions")

	flag.Parse()

	questions := csvRead.ReadCSV(params.file)

	if params.shuffle {
		questions = shuffleQuestions(questions)
	}

	fmt.Println("Welcome to the Quiz Game!")

	timer := time.NewTimer(time.Duration(params.quizTime) * time.Second)

	corr, quiz_len := quizGame(questions, timer)

	fmt.Printf("You have scored %d out of %d questions", corr, quiz_len)

}
