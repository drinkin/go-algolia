package algolia_test

import (
	"github.com/drinkin/di/random"
	"github.com/drinkin/go-algolia/algolia"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Algolia", func() {
	var (
		example *Example

		clients = []algolia.Client{
			algolia.FromEnv(true), // mock client
			algolia.FromEnv(),     // real client
		}
	)

	BeforeEach(func() {
		example = RandomExample()
	})

	CheckValidIndex := func(idx algolia.Index) {
		It("success", func() {
			Expect(idx.Name()).To(Equal(TestIndexName))
			tr, err := idx.UpdateObject(example)
			Expect(err).ToNot(HaveOccurred())
			Expect(tr.ObjectId).To(Equal(example.AlgoliaId()))

			err = tr.Wait()
			Expect(err).ToNot(HaveOccurred())

			savedObj := new(Example)
			Expect(idx.GetObject(example.AlgoliaId()).Scan(savedObj)).ToNot(HaveOccurred())
			Expect(savedObj).To(Equal(example))

			example2 := RandomExample()
			bt, err := idx.BatchUpdate([]algolia.Indexable{example2})
			Expect(err).ToNot(HaveOccurred())
			Expect(example2.AlgoliaId()).To(Equal(bt.ObjectIds[0]))
		})

		It("Get object that doesn't exist", func() {
			obj := new(Example)
			err := idx.GetObject(random.Base64(10)).Scan(obj)
			Expect(err).To(HaveOccurred())
		})

		It("Delete object that doesn't exist", func() {
			obj := new(Example)
			err := idx.DeleteObject(random.Base64(10)).Scan(obj)
			Expect(err).To(HaveOccurred())
		})

		It("GetTaskStatus that doesn't exist", func() {
			ts, err := idx.GetTaskStatus(random.Int64(1, 99999999999))
			Expect(err).ToNot(HaveOccurred())
			Expect(ts.Status).To(Equal("notPublished"))
		})

		It("Can set settings", func() {
			idx.Settings().
				AttributesToIndex("name").
				AttributesForFaceting("facet_1", "facet_2").
				CustomRanking("desc(name)").
				Save()
		})

		It("can clear", func() {
			tr, err := idx.Clear()
			Expect(err).ToNot(HaveOccurred())
			err = tr.Wait()
			Expect(err).ToNot(HaveOccurred())
		})
	}

	CheckValidClient := func(client algolia.Client) {
		CheckValidIndex(client.Index(TestIndexName))

		It("can set prefix", func() {
			client.SetIndexPrefix("test_")
			Expect(client.Index("hello").Name()).To(Equal("test_hello"))
		})
	}

	for _, client := range clients {
		CheckValidClient(client)
	}

})
