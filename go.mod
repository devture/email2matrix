module devture-email2matrix

go 1.12

require (
	github.com/asaskevich/EventBus v0.0.0-20180315140547-d46933a94f05 // indirect
	github.com/euskadi31/go-service v1.3.1
	github.com/flashmob/go-guerrilla v0.0.0-20190620040312-f9c49656c2db
	github.com/go-sql-driver/mysql v1.4.1 // indirect
	github.com/jhillyerd/enmime v0.5.0
	github.com/matrix-org/gomatrix v0.0.0-20190528120928-7df988a63f26
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/sirupsen/logrus v1.4.2
	gopkg.in/russross/blackfriday.v2 v2.0.0-00010101000000-000000000000
)

replace gopkg.in/russross/blackfriday.v2 => github.com/russross/blackfriday/v2 v2.0.1
