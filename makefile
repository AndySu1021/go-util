mock:
	mockgen -source interface/http.go -destination mock/http_mock.go -package mock
	mockgen -source interface/redis.go -destination mock/redis_mock.go -package mock
	mockgen -source interface/repo.go -destination mock/repo_mock.go -package mock
	mockgen -source interface/etcd.go -destination mock/etcd_mock.go -package mock
	mockgen -source interface/storage.go -destination mock/storage_mock.go -package mock