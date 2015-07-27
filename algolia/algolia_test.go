package algolia_test

import (
	"github.com/drinkin/go-algolia/algolia"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Algolia", func() {
	var (
		example *Example
		client  *algolia.Client
		idx     algolia.Index
	)

	BeforeEach(func() {
		example = RandomExample()
		client = algolia.FromEnv()
		idx = client.Index(TestIndexName)
	})

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
})
