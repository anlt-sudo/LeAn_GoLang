package service

import (
	"io/ioutil"
	"net/http"
	"strconv"
)

type Job struct {
	ID   int
	Data string
}

type Result struct {
	JobID  int
	Output string
	Error  error
}

func FetchDataFromURL(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return "Data from " + url + " has length " + strconv.Itoa(len(body)), nil
}

func Worker(id int, jobs <-chan Job, results chan<- Result) {
	for job := range jobs {
		data, err := FetchDataFromURL(job.Data)
		results <- Result{JobID: job.ID, Output: data, Error: err}
	}
}

func ShowResult(results <-chan Result) {
	for result := range results {
		if result.Error != nil {
			println("Job", result.JobID, "failed:", result.Error.Error())
		} else {
			println("Job", result.JobID, "completed:", result.Output)
		}
	}
}

