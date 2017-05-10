package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"

	"errors"

	"github.com/golang/mock/gomock"
)

func TestCautiousSwimmer(t *testing.T) {
	RegisterFailHandler(Fail)

	RunSpecs(t, "CautiousSwimmer Suite")
}

var _ = Describe("CautiousSwimmer", func() {
	// declare all the vars that will be be needed throughout the test
	var (
		subject       *CautiousSwimmer      // The swimmer
		thermometer   *MockWaterThermometer // magic thing that can measure water temperature by place's name
		place         string                // where are we e.g. Frishman Beach
		okTemperature int                   // If water is colder than this, the swimmer won't swim
		mockCtrl      *gomock.Controller    // this is needed to go-mock
		err           error
	)

	// in every example we'll need to to have fresh instance of thermometer mock
	// could be placed in JustBeforeEach below, but left here for example purposes
	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		thermometer = NewMockWaterThermometer(mockCtrl)
	})

	// this will assure that all the registered calls have happened
	AfterEach(func() {
		mockCtrl.Finish()
	})

	// Separating Creation and Configuration
	JustBeforeEach(func() {
		subject = NewCautiousSwimmer(place, okTemperature, thermometer)
		err = subject.Swim()
	})

	Describe("Swim", func() {
		Context("Thermometer works", func() {
			// Configuration
			BeforeEach(func() {
				place = "bath"
				okTemperature = 26
			})

			Context("Warm water", func() {
				// Setting expectation on mocked thermometer
				BeforeEach(func() {
					thermometer.EXPECT().MeasureAt(place).Return(30, nil)
				})

				// checking for errors
				It("does not complain", func() {
					Expect(err).NotTo(HaveOccurred())
				})

				// checking values
				It("is swimming", func() {
					Expect(subject.State()).To(Equal(StateSwimmin))
				})
			})

			Context("OK temperature water", func() {
				BeforeEach(func() {
					// Calls with any arguments
					// Setting call proxy
					// Logging
					thermometer.EXPECT().MeasureAt(gomock.Any()).Do(func(key string) {
						GinkgoT().Log("Called thermometer with params:", key)
					}).Return(okTemperature, nil)
				})

				// not very interesting
				It("does not complain", func() {
					Expect(err).NotTo(HaveOccurred())
				})

				It("is swimming", func() {
					Expect(subject.State()).To(Equal(StateSwimmin))
				})
			})

			Context("Cold water", func() {
				BeforeEach(func() {
					// Any call that has happened is removed from call lookup map
					// See: Times(), MinTimes(), MaxTimes(), AnyTimes()
					thermometer.EXPECT().MeasureAt(gomock.Any()).Return(20, nil).Times(1)
				})

				// not very interesting
				It("does not complain", func() {
					Expect(err).NotTo(HaveOccurred())
				})

				It("is chillin", func() {
					Expect(subject.State()).To(Equal(StateChillin))
				})
			})
		})

		Context("Thermometer is broken", func() {
			BeforeEach(func() {
				place = "swimming pool"
				okTemperature = 26

				thermometer.EXPECT().MeasureAt(gomock.Any()).Return(0, errors.New("Too lazy to measure"))
			})

			// Omega (alt+z on mac) if you're feeling cool - same as Expect()
			// Should() is the same as To()
			It("complains", func() {
				Ω(err).Should(HaveOccurred())
			})

			// Equal() uses reflect.DeepEqual, this will use strict ==
			It("does not know what is it doing", func() {
				Ω(subject.State()).Should(BeIdenticalTo(StateUnknown))
			})
		})
	})
})
