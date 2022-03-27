package service

import (
	"example/golang-gin-poc/entity"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const (
	TITLE       = "Video Title"
	DESCRIPTION = "Video Description"
	URL         = "https://youtu.be/JqW-i2QjgHQ"
	FIRSTNAME   = "John"
	LASTNAME    = "Doe"
	EMAIL       = "jdoe@mail.com"
)

var testVideo entity.Video = entity.Video{
	Title:       TITLE,
	Description: DESCRIPTION,
	URL:         URL,
	Author: entity.Person{
		FirstName: FIRSTNAME,
		LastName:  LASTNAME,
		Email:     EMAIL,
	},
}

var testVideoService VideoService

var _ = BeforeSuite(func() {
	testVideoService = New()
})

var _ = AfterSuite(func() {
})

var _ = Describe("Video Service", func() {
	Describe("Fetching all existing videos", func() {
		Context("If there is a video in the database", func() {
			BeforeEach(func() {
				testVideoService.Save(testVideo)
			})

			It("should return at least one element", func() {
				videoList := testVideoService.FindAll()

				Ω(videoList).ShouldNot(BeEmpty())
			})
			It("should map the fields correctly", func() {
				firstVideo := testVideoService.FindAll()[0]

				Ω(firstVideo.Title).Should(Equal(TITLE))
				Ω(firstVideo.Description).Should(Equal(DESCRIPTION))
				Ω(firstVideo.URL).Should(Equal(URL))
				Ω(firstVideo.Author.FirstName).Should(Equal(FIRSTNAME))
				Ω(firstVideo.Author.LastName).Should(Equal(LASTNAME))
				Ω(firstVideo.Author.Email).Should(Equal(EMAIL))
			})

			AfterEach(func() {
				video := testVideoService.FindAll()[0]
				testVideoService.Delete(video)
			})
		})

		Context("If there are no videos in the database", func() {
			It("should return an empty list", func() {
				videos := testVideoService.FindAll()
				Ω(videos).Should(BeEmpty())
			})
		})
	})
})
