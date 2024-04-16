package s3

import (
	"cloud.google.com/go/storage"
)

type Storage struct {
	Client     *storage.Client
	ProjectId  string
	BucketName string
}

// func (s *Storage) Init() {
// 	creds, err := google.FindDefaultCredentials(context.Background())
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}

// 	client, err := storage.NewClient(
// 		context.Background(),
// 		option.WithCredentialsJSON(creds.JSON),
// 	)
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}

// 	s.Client = client
// 	s.ProjectId = os.Getenv("GOOGLE_PROJECT_ID")
// 	s.BucketName = "terraform-dep"
// }

// func (s *Storage) Upload(objectName string) {
// 	wc := s.Client.Bucket(s.BucketName).Object("").NewWriter(context.Background())
// 	fmt.Println(wc)
// }
