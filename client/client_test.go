package client_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	devsClient "api-hackaton-devs/client"

	"github.com/go-kit/log"
	"github.com/stretchr/testify/assert"
)

var (
	timeoutDuration = time.Duration(5) * time.Second
	logg            = log.NewNopLogger()
)

func TestClient(t *testing.T) {
	t.Run("Success client execution", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(clientResponseJSON))
		}))
		defer server.Close()
		client := devsClient.NewDevsClient(server.URL, timeoutDuration, logg)
		res, err := client.Call(context.TODO(), &devsClient.Request{})
		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.NotEmpty(t, res.Results)
		assert.Equal(t, len(res.Results), 10)
	})

	t.Run("Fails connection - wrong URL", func(t *testing.T) {
		client := devsClient.NewDevsClient("asljkdaskldb", timeoutDuration, logg)
		res, err := client.Call(context.TODO(), &devsClient.Request{})
		assert.NotNil(t, err)
		assert.Nil(t, res)
	})

	t.Run("Create a new client", func(t *testing.T) {
		client := devsClient.NewDevsClient("", timeoutDuration, logg)
		assert.NotNil(t, client)
		assert.Implements(t, (*devsClient.DevsClient)(nil), client)
	})
}

var (
	clientResponseJSON = `
	{
		"results": [
			{
				"gender": "male",
				"name": {
					"title": "Mr",
					"first": "آرتين",
					"last": "سالاری"
				},
				"location": {
					"street": {
						"number": 6182,
						"name": "مجاهدین اسلام"
					},
					"city": "تهران",
					"state": "خراسان جنوبی",
					"country": "Iran",
					"postcode": 39508,
					"coordinates": {
						"latitude": "33.3278",
						"longitude": "-39.4874"
					},
					"timezone": {
						"offset": "-1:00",
						"description": "Azores, Cape Verde Islands"
					}
				},
				"email": "artyn.slry@example.com",
				"login": {
					"uuid": "04699545-d8d3-43ab-a766-9314245f4a3e",
					"username": "greengorilla650",
					"password": "intruder",
					"salt": "2FZpnNRx",
					"md5": "1126bfd7f9062490e1071d5c5d217faa",
					"sha1": "deef869cbcead9f5a96a1c93f3cb61aa440ba2c5",
					"sha256": "b439dc434446b9220c1bc909c0de1685b7197805e4c556cf98853a7789ecf8af"
				},
				"dob": {
					"date": "1955-02-06T21:17:42.106Z",
					"age": 67
				},
				"registered": {
					"date": "2004-05-02T02:47:25.039Z",
					"age": 18
				},
				"phone": "012-39362910",
				"cell": "0963-321-3873",
				"id": {
					"name": "",
					"value": null
				},
				"picture": {
					"large": "https://randomuser.me/api/portraits/men/27.jpg",
					"medium": "https://randomuser.me/api/portraits/med/men/27.jpg",
					"thumbnail": "https://randomuser.me/api/portraits/thumb/men/27.jpg"
				},
				"nat": "IR"
			},
			{
				"gender": "male",
				"name": {
					"title": "Mr",
					"first": "Henning",
					"last": "Thorkildsen"
				},
				"location": {
					"street": {
						"number": 8963,
						"name": "Ole Wingers vei"
					},
					"city": "Bergen",
					"state": "Telemark",
					"country": "Norway",
					"postcode": "0691",
					"coordinates": {
						"latitude": "-58.2755",
						"longitude": "-146.2421"
					},
					"timezone": {
						"offset": "-5:00",
						"description": "Eastern Time (US & Canada), Bogota, Lima"
					}
				},
				"email": "henning.thorkildsen@example.com",
				"login": {
					"uuid": "a56325db-6c45-4d3e-a0c5-d9c471f8f8e3",
					"username": "orangeswan274",
					"password": "zhuang",
					"salt": "YJWqCh7I",
					"md5": "e208f4bfb0467d3f847a88b85b0217af",
					"sha1": "8e54710877ecdd1503f41d7e0bdab9ca4e31a5f0",
					"sha256": "2d9ac0d9fc1bc1fddd69f3431a2eba935e42d3bb3400898936002e8e5b8ba3df"
				},
				"dob": {
					"date": "1967-05-27T18:43:39.209Z",
					"age": 55
				},
				"registered": {
					"date": "2018-09-09T02:49:10.394Z",
					"age": 4
				},
				"phone": "34251476",
				"cell": "45336558",
				"id": {
					"name": "FN",
					"value": "27056742366"
				},
				"picture": {
					"large": "https://randomuser.me/api/portraits/men/15.jpg",
					"medium": "https://randomuser.me/api/portraits/med/men/15.jpg",
					"thumbnail": "https://randomuser.me/api/portraits/thumb/men/15.jpg"
				},
				"nat": "NO"
			},
			{
				"gender": "male",
				"name": {
					"title": "Mr",
					"first": "محمدامين",
					"last": "احمدی"
				},
				"location": {
					"street": {
						"number": 5084,
						"name": "آذربایجان"
					},
					"city": "کاشان",
					"state": "گلستان",
					"country": "Iran",
					"postcode": 94009,
					"coordinates": {
						"latitude": "-14.2981",
						"longitude": "158.5157"
					},
					"timezone": {
						"offset": "-10:00",
						"description": "Hawaii"
					}
				},
				"email": "mhmdmyn.hmdy@example.com",
				"login": {
					"uuid": "2364817b-a295-45a8-9e21-b6b0180eee46",
					"username": "brownrabbit295",
					"password": "cheers",
					"salt": "sz7NRaHO",
					"md5": "6f5527e5b03c65465f07a205ef51ec73",
					"sha1": "cdfd46b466bcb7e2c3604f88fefbebd516e902ff",
					"sha256": "499858179cb8b041771a4845db47cbd385cf18cbe53749ff3fa655892b3277fd"
				},
				"dob": {
					"date": "1994-10-11T18:41:11.883Z",
					"age": 28
				},
				"registered": {
					"date": "2011-05-16T02:56:55.541Z",
					"age": 11
				},
				"phone": "053-13337648",
				"cell": "0974-254-9194",
				"id": {
					"name": "",
					"value": null
				},
				"picture": {
					"large": "https://randomuser.me/api/portraits/men/99.jpg",
					"medium": "https://randomuser.me/api/portraits/med/men/99.jpg",
					"thumbnail": "https://randomuser.me/api/portraits/thumb/men/99.jpg"
				},
				"nat": "IR"
			},
			{
				"gender": "female",
				"name": {
					"title": "Miss",
					"first": "Florence",
					"last": "Holmes"
				},
				"location": {
					"street": {
						"number": 754,
						"name": "Hunters Creek Dr"
					},
					"city": "Sydney",
					"state": "South Australia",
					"country": "Australia",
					"postcode": 5430,
					"coordinates": {
						"latitude": "78.8035",
						"longitude": "49.9910"
					},
					"timezone": {
						"offset": "+4:00",
						"description": "Abu Dhabi, Muscat, Baku, Tbilisi"
					}
				},
				"email": "florence.holmes@example.com",
				"login": {
					"uuid": "b08a33c3-e8f8-4525-a5bb-6dc7387f2e70",
					"username": "bigleopard487",
					"password": "asgard",
					"salt": "OpE0WxUv",
					"md5": "6f7d6e9996379e2202f21866c6d68dd7",
					"sha1": "da536609d3ec88f30edd4d06feb800b24f594392",
					"sha256": "309d18bb3a225bd22fd107a32471df89196eebf3173d95abe9dac36d8ef7e48e"
				},
				"dob": {
					"date": "1992-03-12T06:35:30.407Z",
					"age": 30
				},
				"registered": {
					"date": "2010-09-01T09:54:13.413Z",
					"age": 12
				},
				"phone": "07-5107-9892",
				"cell": "0427-512-319",
				"id": {
					"name": "TFN",
					"value": "833524441"
				},
				"picture": {
					"large": "https://randomuser.me/api/portraits/women/35.jpg",
					"medium": "https://randomuser.me/api/portraits/med/women/35.jpg",
					"thumbnail": "https://randomuser.me/api/portraits/thumb/women/35.jpg"
				},
				"nat": "AU"
			},
			{
				"gender": "male",
				"name": {
					"title": "Mr",
					"first": "Veit",
					"last": "Lampert"
				},
				"location": {
					"street": {
						"number": 6724,
						"name": "Uhlandstraße"
					},
					"city": "Zirndorf",
					"state": "Sachsen-Anhalt",
					"country": "Germany",
					"postcode": 63193,
					"coordinates": {
						"latitude": "17.1681",
						"longitude": "-57.6007"
					},
					"timezone": {
						"offset": "-3:30",
						"description": "Newfoundland"
					}
				},
				"email": "veit.lampert@example.com",
				"login": {
					"uuid": "62a16c96-a37f-4c6b-9118-6e75b920cfde",
					"username": "brownsnake556",
					"password": "kennedy",
					"salt": "p2odCbG7",
					"md5": "44f48115fafbc7d299c9f1a633e9ed6c",
					"sha1": "3d2c86152bee080710723fe7ba158b3fb1b00885",
					"sha256": "cd005b6d369880f8c792133ad642bc5e5511a72a4437338b446f87258fa1e923"
				},
				"dob": {
					"date": "1995-03-19T08:38:39.386Z",
					"age": 27
				},
				"registered": {
					"date": "2011-10-03T22:03:12.226Z",
					"age": 11
				},
				"phone": "0307-6414570",
				"cell": "0172-6059167",
				"id": {
					"name": "",
					"value": null
				},
				"picture": {
					"large": "https://randomuser.me/api/portraits/men/24.jpg",
					"medium": "https://randomuser.me/api/portraits/med/men/24.jpg",
					"thumbnail": "https://randomuser.me/api/portraits/thumb/men/24.jpg"
				},
				"nat": "DE"
			},
			{
				"gender": "female",
				"name": {
					"title": "Miss",
					"first": "Flory",
					"last": "Ten Haaf"
				},
				"location": {
					"street": {
						"number": 2199,
						"name": "Keenstraat"
					},
					"city": "Damwald",
					"state": "Drenthe",
					"country": "Netherlands",
					"postcode": 18217,
					"coordinates": {
						"latitude": "-44.0267",
						"longitude": "-94.6417"
					},
					"timezone": {
						"offset": "0:00",
						"description": "Western Europe Time, London, Lisbon, Casablanca"
					}
				},
				"email": "flory.tenhaaf@example.com",
				"login": {
					"uuid": "b91452ec-4484-4667-a840-f84aefd94d30",
					"username": "ticklishtiger289",
					"password": "howdy",
					"salt": "knTDNGpd",
					"md5": "65a33769a10a06dcdf5668c23417b6b8",
					"sha1": "3a074325e80f6f0a037b4fdbc02ca79f3b6d7b31",
					"sha256": "1506d3513248446c42e829d71452217423b129f05c82ba4bfb61a2beefe1b8f0"
				},
				"dob": {
					"date": "1958-02-13T21:21:25.338Z",
					"age": 64
				},
				"registered": {
					"date": "2018-05-12T17:12:30.003Z",
					"age": 4
				},
				"phone": "(508)-670-0592",
				"cell": "(724)-672-8635",
				"id": {
					"name": "BSN",
					"value": "63828611"
				},
				"picture": {
					"large": "https://randomuser.me/api/portraits/women/36.jpg",
					"medium": "https://randomuser.me/api/portraits/med/women/36.jpg",
					"thumbnail": "https://randomuser.me/api/portraits/thumb/women/36.jpg"
				},
				"nat": "NL"
			},
			{
				"gender": "female",
				"name": {
					"title": "Ms",
					"first": "Evamaria",
					"last": "Gabriel"
				},
				"location": {
					"street": {
						"number": 6704,
						"name": "Schulstraße"
					},
					"city": "Fladungen",
					"state": "Baden-Württemberg",
					"country": "Germany",
					"postcode": 76108,
					"coordinates": {
						"latitude": "-7.4817",
						"longitude": "-171.8804"
					},
					"timezone": {
						"offset": "0:00",
						"description": "Western Europe Time, London, Lisbon, Casablanca"
					}
				},
				"email": "evamaria.gabriel@example.com",
				"login": {
					"uuid": "8f83d8e8-360c-4b72-b9b1-4a08e1acc94b",
					"username": "redpanda297",
					"password": "biker",
					"salt": "quqx6GCA",
					"md5": "9ff558e1dcd5fa9a097335efba80f388",
					"sha1": "50a8ee6941ebf98d5427a526e185d3ef7c560b7a",
					"sha256": "be909aaa44281d16aa909f20a606d12e8e60eaec81a33eeae7abbb261030bb51"
				},
				"dob": {
					"date": "1996-07-22T13:45:45.779Z",
					"age": 26
				},
				"registered": {
					"date": "2016-08-26T07:57:12.949Z",
					"age": 6
				},
				"phone": "0663-7317852",
				"cell": "0174-0693440",
				"id": {
					"name": "",
					"value": null
				},
				"picture": {
					"large": "https://randomuser.me/api/portraits/women/80.jpg",
					"medium": "https://randomuser.me/api/portraits/med/women/80.jpg",
					"thumbnail": "https://randomuser.me/api/portraits/thumb/women/80.jpg"
				},
				"nat": "DE"
			},
			{
				"gender": "male",
				"name": {
					"title": "Mr",
					"first": "Jacob",
					"last": "Ross"
				},
				"location": {
					"street": {
						"number": 307,
						"name": "Dalhousie Ave"
					},
					"city": "Lumsden",
					"state": "Alberta",
					"country": "Canada",
					"postcode": "P6R 3J7",
					"coordinates": {
						"latitude": "-80.0087",
						"longitude": "-13.0487"
					},
					"timezone": {
						"offset": "0:00",
						"description": "Western Europe Time, London, Lisbon, Casablanca"
					}
				},
				"email": "jacob.ross@example.com",
				"login": {
					"uuid": "3e20123b-7a84-407d-a33e-681d9b008e4c",
					"username": "crazybear532",
					"password": "donovan",
					"salt": "z7rQlH0x",
					"md5": "3b4cc15098f165407aa5d019d9bd837b",
					"sha1": "f2f98a2062f7fe6276b2ccba65c20d3c4e79599d",
					"sha256": "1f6bca41b56edbebf62e09cf77e8ae6a57c5ac5617eb08cdbb3ef4d2d6075965"
				},
				"dob": {
					"date": "1974-06-04T22:07:45.062Z",
					"age": 48
				},
				"registered": {
					"date": "2004-11-07T02:31:00.391Z",
					"age": 18
				},
				"phone": "957-129-6432",
				"cell": "487-419-4774",
				"id": {
					"name": "",
					"value": null
				},
				"picture": {
					"large": "https://randomuser.me/api/portraits/men/29.jpg",
					"medium": "https://randomuser.me/api/portraits/med/men/29.jpg",
					"thumbnail": "https://randomuser.me/api/portraits/thumb/men/29.jpg"
				},
				"nat": "CA"
			},
			{
				"gender": "female",
				"name": {
					"title": "Mrs",
					"first": "Sedef",
					"last": "Nebioğlu"
				},
				"location": {
					"street": {
						"number": 7191,
						"name": "Şehitler Cd"
					},
					"city": "Bursa",
					"state": "Erzincan",
					"country": "Turkey",
					"postcode": 52978,
					"coordinates": {
						"latitude": "-40.8761",
						"longitude": "-5.4590"
					},
					"timezone": {
						"offset": "+7:00",
						"description": "Bangkok, Hanoi, Jakarta"
					}
				},
				"email": "sedef.nebioglu@example.com",
				"login": {
					"uuid": "bb3ea68a-5e4d-4a22-8843-f7ad40b4546e",
					"username": "lazykoala663",
					"password": "jones1",
					"salt": "oOaTxw0q",
					"md5": "fa8aa3e67d51f82b5c06cde15824b9f1",
					"sha1": "0730a434b0a56e2d0028ecf47f914ea219b7fda3",
					"sha256": "600607bc54907b337553de41fb5bfa19dd81a50236e7ae662c4eeea9cd9cf35c"
				},
				"dob": {
					"date": "1988-05-20T03:23:54.903Z",
					"age": 34
				},
				"registered": {
					"date": "2016-05-24T04:33:55.660Z",
					"age": 6
				},
				"phone": "(980)-241-1788",
				"cell": "(218)-868-7496",
				"id": {
					"name": "",
					"value": null
				},
				"picture": {
					"large": "https://randomuser.me/api/portraits/women/65.jpg",
					"medium": "https://randomuser.me/api/portraits/med/women/65.jpg",
					"thumbnail": "https://randomuser.me/api/portraits/thumb/women/65.jpg"
				},
				"nat": "TR"
			},
			{
				"gender": "female",
				"name": {
					"title": "Ms",
					"first": "Firmina",
					"last": "Almeida"
				},
				"location": {
					"street": {
						"number": 3350,
						"name": "Rua Onze "
					},
					"city": "Salvador",
					"state": "Distrito Federal",
					"country": "Brazil",
					"postcode": 70236,
					"coordinates": {
						"latitude": "4.7174",
						"longitude": "-12.1633"
					},
					"timezone": {
						"offset": "+4:00",
						"description": "Abu Dhabi, Muscat, Baku, Tbilisi"
					}
				},
				"email": "firmina.almeida@example.com",
				"login": {
					"uuid": "62ab3fdd-a95c-452b-a847-1ae5770267c7",
					"username": "sadfish640",
					"password": "skinner",
					"salt": "NCfmAxVz",
					"md5": "33d73483dbc968ba8098c0bbc545e7cc",
					"sha1": "33cad6637d3b4f2791d743597b997e95080d98fd",
					"sha256": "19830eb011ee20a1758bcfb98bcd2c6e81e39c1cdec71829c7a6bb02ad875a66"
				},
				"dob": {
					"date": "1988-11-21T02:30:21.633Z",
					"age": 34
				},
				"registered": {
					"date": "2006-03-05T08:35:58.515Z",
					"age": 16
				},
				"phone": "(81) 4133-0488",
				"cell": "(44) 5217-6598",
				"id": {
					"name": "",
					"value": null
				},
				"picture": {
					"large": "https://randomuser.me/api/portraits/women/5.jpg",
					"medium": "https://randomuser.me/api/portraits/med/women/5.jpg",
					"thumbnail": "https://randomuser.me/api/portraits/thumb/women/5.jpg"
				},
				"nat": "BR"
			}
		],
		"info": {
			"seed": "dcfcaef78482b443",
			"results": 10,
			"page": 1,
			"version": "1.3"
		}
	}
	`
)
