package algolia_test

import (
	"github.com/drinkin/di/random"
	"github.com/drinkin/go-algolia/algolia"
	"github.com/k0kubun/pp"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Algolia", func() {
	var (
		example *Example
	)

	BeforeEach(func() {
		example = RandomExample()
	})

	CheckValidIndex := func(idx algolia.Index) {
		expectTaskPublished := func(taskId int64) {
			Eventually(func() bool {
				return idx.Must().GetTaskStatus(taskId).IsPublished()
			}, 1, .1).Should(BeTrue())
		}

		It("", func() {
			tr, err := idx.UpdateObject(example)
			Expect(err).ToNot(HaveOccurred())
			Expect(tr.ObjectId).To(Equal(example.AlgoliaId()))

			expectTaskPublished(tr.TaskId)

			savedObj := new(Example)
			Expect(idx.GetObject(example.AlgoliaId()).Scan(savedObj)).ToNot(HaveOccurred())
			Expect(savedObj).To(Equal(example))

			example2 := RandomExample()
			bt := idx.Must().BatchUpdate([]algolia.Indexable{example2})
			Expect(example2.AlgoliaId()).To(Equal(bt.ObjectIds[0]))
		})

		It("Get object that doesn't exist", func() {
			obj := new(Example)
			err := idx.GetObject(random.Base64(10)).Scan(obj)
			Expect(err).To(HaveOccurred())
		})

		It("GetTaskStatus that doesn't exist", func() {
			ts, err := idx.GetTaskStatus(random.Int64(1, 99999999999))
			pp.Print(ts)
			pp.Print(err)
		})

	}
	CheckValidIndex(algolia.FromEnv().MockIndex(TestIndexName))
	CheckValidIndex(algolia.FromEnv().Index(TestIndexName))
})
