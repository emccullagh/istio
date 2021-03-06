//  Copyright 2019 Istio Authors
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package util

const (
	// Server certificate, private key and CA certificate
	TLSServerCertA = `-----BEGIN CERTIFICATE-----
MIIFPzCCAyegAwIBAgIDEAISMA0GCSqGSIb3DQEBCwUAMEQxCzAJBgNVBAYTAlVT
MQ8wDQYDVQQIDAZEZW5pYWwxDDAKBgNVBAoMA0RpczEWMBQGA1UEAwwNKi5leGFt
cGxlLmNvbTAeFw0xOTA1MjkyMzEzMjFaFw0yOTA1MjYyMzEzMjFaMFoxCzAJBgNV
BAYTAlVTMQ8wDQYDVQQIDAZEZW5pYWwxFDASBgNVBAcMC1NwcmluZ2ZpZWxkMQww
CgYDVQQKDANEaXMxFjAUBgNVBAMMDSouZXhhbXBsZS5jb20wggEiMA0GCSqGSIb3
DQEBAQUAA4IBDwAwggEKAoIBAQC7M6icT8vcZ3KZpTZW3zOUGBpJCRG5HKfclyZ7
ONY9Dnc/3+nESJ0vqIOnV7/NbqjWoJYQ8v1xoUPRvAKA6nVSXmvHEfOuq0LofYFP
DoC0o6WWi6VslZEOtf87jzcGLbLbBkBLbmd5/LWp5zw4Bexe0YHKSnsObcCphleN
DeCaPA73Z9YzDEL53PLyH0emuXKyx+lsDcijv+ualVu4HACpeo354Df3XzyxH3TT
5+sLLBVf4ZsyFsP/VtZYwP9PCWepgPH0l1ekPN7AE8MQQmddJUjTFepnRaVMvwmt
dWirp/S1TVYawGLIN/SlX+BHvKPvcd0GXxKm8AjdHD2q7ZxtAgMBAAGjggEiMIIB
HjAJBgNVHRMEAjAAMBEGCWCGSAGG+EIBAQQEAwIGQDAzBglghkgBhvhCAQ0EJhYk
T3BlblNTTCBHZW5lcmF0ZWQgU2VydmVyIENlcnRpZmljYXRlMB0GA1UdDgQWBBQW
ljUOEibJ9/dJ8dw1UOV8Wcg18TCBhAYDVR0jBH0we4AU44b/zezM4Uz3S6WXtXN+
cQoj46WhXqRcMFoxCzAJBgNVBAYTAlVTMQ8wDQYDVQQIDAZEZW5pYWwxFDASBgNV
BAcMC1NwcmluZ2ZpZWxkMQwwCgYDVQQKDANEaXMxFjAUBgNVBAMMDSouZXhhbXBs
ZS5jb22CAxACEjAOBgNVHQ8BAf8EBAMCBaAwEwYDVR0lBAwwCgYIKwYBBQUHAwEw
DQYJKoZIhvcNAQELBQADggIBABpv2bmPbl3meDpLQ2cSH9QoDitkYBMxwp7cL6GA
4Orv6SPbE+PIiijpSSMKbbtMgliL8eJwxAjyPUms0djtH2hvDLZXnRmCVHm4+SOy
CSU7PofELigd1B9w9BRQ183AWzClvtKoS/fVq8szHtLVy+5rgj6bigskdis6lnyN
NzcXusi2Rp+BmcrNRsfin223dbrEG5qeJqTOniEtHyrIzFUYUK9RGtdSVq/k4+0m
DQuG9yK7oqrhL5aSzIs89RD5ofDadhBZJx6OukIppO3aqDJPt8JKI72hhaJc0YNG
CgbluqC2CCZK0BA9JB9c4WGDwr+yvmXxXTO/z5Tk27Pac3vVyEGV4Se5bG1ZT5it
ctsLyKXW/GMWe55T3A5u/0ebvIM7VRquGq4khFq/Rm5uRJhkYBG2j40SHSoIZGV7
bsQR2ucH//EOyWHN96JGrkKV9QaeK/+kV0u7rvqTbzMDKjp2dH0UxTEE7ToHczam
iNNYj8XTzQ7BsqQY5Ti1N2MKxfzgtmik+3BEP42pr9RS9d4PCE/SxF6DOEcP9nZc
5FbxAztiqK6WRZoXegLKEYylA/kT6BKk/M42jJ7SWH76f+Breylo28w+XRX01DPu
3YtkgFZMmdf6fdZLgo3uOUP2aBUZYneHNMDDmt+/VqyykwpHLBeeJ4cAXikmWQQY
rMx/
-----END CERTIFICATE-----`
	TLSServerKeyA = `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAuzOonE/L3GdymaU2Vt8zlBgaSQkRuRyn3JcmezjWPQ53P9/p
xEidL6iDp1e/zW6o1qCWEPL9caFD0bwCgOp1Ul5rxxHzrqtC6H2BTw6AtKOlloul
bJWRDrX/O483Bi2y2wZAS25nefy1qec8OAXsXtGBykp7Dm3AqYZXjQ3gmjwO92fW
MwxC+dzy8h9HprlyssfpbA3Io7/rmpVbuBwAqXqN+eA39188sR900+frCywVX+Gb
MhbD/1bWWMD/TwlnqYDx9JdXpDzewBPDEEJnXSVI0xXqZ0WlTL8JrXVoq6f0tU1W
GsBiyDf0pV/gR7yj73HdBl8SpvAI3Rw9qu2cbQIDAQABAoIBAFtQd47H3clFLMGF
tVvqxF4Y47l8kwiY0cjocfzpXJer1r9xmbYFNadpq86VHxo2QGVxL8JUxQwIfexw
qUf5FgF3zVrthtBM+fYuQkUt7doveTZWJ/svefKTFE0ZLt30rKnqArLIx/DZNEQM
Z3NscEBQMd9bNWt0XeZTgaFsg4K3sfoSNm3ksQU57MHB9GGWEyjYOLuGc3qKMVnP
MJ71Ks9OUH/6t9yNGuKKuwt/0PV/rrjbz2rPgsxyHeiRLftNjP6CibLeKbDBDisJ
3Gmt7ehqXu6PusOGIMZQGEgjf8FangjW/AmutJos6/18FcTPAF3syroIg8pwvp+5
CSJDtIECgYEA9sqZ2kBq9SL6r+oLppP3FOSfDA5faVwAX/KEXzvNzWMYDq5Im9b6
maxnvjvkh8Bp2whBfh/ny5JX7xJ+/bnjjSb5SsOATnoBx0pFfayXmjhe/lYWQDdh
0JMZZKSZ6b2j94cbCwgqmOqvH7pm1RoGw3TfQoL3YiBHxoyqIDf1P3cCgYEAwi/Y
nWNsAu7iX8ROpahURDA28jiH0EgF67h2YVZVTBiR5/1YH6fps85YqVjCU773PTiI
ulv+bdLRdpBzihE2MAC9ZrM0iGelclWUz1rfJR3q4HRswVoqRFpw3fxVqrXwxlBd
0Vh1AGB4ICyCf8SbcOdgTtzwzdf7eECf6fqu5DsCgYBpAYl2+MPJq+l4YiA0725J
IIGf0pm8LZ4nNPwnvT1y73/z9yoDzPPx4Q5/PeKd9DvWK2waax7oWks5+Oe5s5dj
nrqhmg+E1JrfHfX1ZDNZNjUNCLoM9sOSyJQcY4DtmHgkwQK7rezaBbkfyeMpNxq7
nYXFOg5iXRPNIlrcvJKIPwKBgCFhnDkiGDVA34Va97Ieh63ZI/jrot05JFtrMF53
Ot6D8sEQmg+HC/Ou6yrhaTaAQ3wwtcBjhA9ZCEz7dAEuCESkO+rlr4grWROYhqZG
lsmIpRHEsdVjLJTnfzOJ5ygbcrKEqjwMmsBf1Qai4CQapiX4gGPOM08Pa9k2MsdE
IsB1AoGBALve0Lwkv8KbL0qkIjso/RO1wq5FP85nrsGA6KRUwhHiI0pEZQV+oeb8
6U6KwFLzTQLNgHPvcDwy0/7UTKYg+N5PzTmP5GwenfaqXsKUXIMBUM9fG4bUN0F0
0aaoHpOg1BKcKVhMOsNFujiyR4l4PN69924bQDbiEROu9Z4IAe14
-----END RSA PRIVATE KEY-----`
	TLSClientCertA = `-----BEGIN CERTIFICATE-----
MIIE4TCCAsmgAwIBAgIDEAITMA0GCSqGSIb3DQEBCwUAMEQxCzAJBgNVBAYTAlVT
MQ8wDQYDVQQIDAZEZW5pYWwxDDAKBgNVBAoMA0RpczEWMBQGA1UEAwwNKi5leGFt
cGxlLmNvbTAeFw0xOTA1MjkyMzEzMjNaFw0yOTA1MjYyMzEzMjNaMFoxCzAJBgNV
BAYTAlVTMQ8wDQYDVQQIDAZEZW5pYWwxFDASBgNVBAcMC1NwcmluZ2ZpZWxkMQww
CgYDVQQKDANEaXMxFjAUBgNVBAMMDSouZXhhbXBsZS5jb20wggEiMA0GCSqGSIb3
DQEBAQUAA4IBDwAwggEKAoIBAQCuvkq8UGWBHgb28aUYCIo/fbB+HcI1RGC3SnqK
skacparseQR4zQJVr1pG/3ANxqeJMg19UVfobdxqj4hp28oEBv3JXxjM6GsnRxkY
7thpFTwgxI9L8XpnQJmj3qf8Zh8R7SVxMLf8+7lfdD4IJ9vggzc3f1YotrwbOEtw
UuQD3g1f0r1JM83t0uA4Jjh8ZTP7PgDZzct37O8PuYV22ADA/yAuJRBmRyt6J7aw
+gI19C7qXAx7r6kI936F2H62n28jZtnXXeQrkFly8bR6qxdAmEzTWL2ojq1ja71M
nFAXDx6T6UOn+Xr2Nl//OWqyo2I05I7l3wT3jshGyWT292cFAgMBAAGjgcUwgcIw
CQYDVR0TBAIwADARBglghkgBhvhCAQEEBAMCBaAwMwYJYIZIAYb4QgENBCYWJE9w
ZW5TU0wgR2VuZXJhdGVkIENsaWVudCBDZXJ0aWZpY2F0ZTAdBgNVHQ4EFgQU4uf4
NTynL454PINH6moFU26G3W0wHwYDVR0jBBgwFoAU44b/zezM4Uz3S6WXtXN+cQoj
46UwDgYDVR0PAQH/BAQDAgXgMB0GA1UdJQQWMBQGCCsGAQUFBwMCBggrBgEFBQcD
BDANBgkqhkiG9w0BAQsFAAOCAgEAkgutVetrMuVeMpfbHWOnr/Wjx6PXkWXchIda
tGXpTCm3/Mr1LTEwZVrdGQSUwSi3cCfd/oF5xpKthvdFYgfQX5ZYvkTOXVnoWQxc
fSzn+nzxjiaOk3WnEw8oB76rd8eVmuaiGuUibEOHw8TzzZCaSSVFKipHftdrG/lD
1fY0+O6cK68Vz/DE4lVRRecO2lmxXAlP6zmohnCJkEgN0panVkUXosalMJ0d+msN
d3UeiqKtEI5MjGoej5HVlIg2UcuNJuaomoURdfF3SFB3PaNSQkXSDlhuoyr+HJyt
7LA/GGT8X3liNy4grRVAb0cSQMcfv9tVZlnR8cKVozw8dDxouJt6Cu895Gx1VDYm
UZTY0l7JOOEktLW8B5DKh+wpiCEkw2ZIfUm8Win0HWghcLAvU287UtNOPsNPQedf
XIWpgWfFTXcYi7x1MiZteg4CHg9C5PjIbhHl8iSxVkxAthQB6YxWc1jXr7kzpEi2
FVmsGr6Czs5KLri6fAqA8WU9ynsGO2eVDV4N1S92eUvsmLLnaR80eZQMIBo6eln9
QgUiVrS1KjpImTngrHgVWfENLxdOdrg3XKQN1pys+52LLDvbSy/h9BUECi3xBxE3
eGPoNAuDveWT5Akxxg3Yy825uA6jNvyOfkyvq+yTbMgFlzU+Wd14uZKLTdtTbuIr
Y7JHxNA=
-----END CERTIFICATE-----`
	TLSClientKeyA = `-----BEGIN RSA PRIVATE KEY-----
MIIEpAIBAAKCAQEArr5KvFBlgR4G9vGlGAiKP32wfh3CNURgt0p6irJGnKWq7HkE
eM0CVa9aRv9wDcaniTINfVFX6G3cao+IadvKBAb9yV8YzOhrJ0cZGO7YaRU8IMSP
S/F6Z0CZo96n/GYfEe0lcTC3/Pu5X3Q+CCfb4IM3N39WKLa8GzhLcFLkA94NX9K9
STPN7dLgOCY4fGUz+z4A2c3Ld+zvD7mFdtgAwP8gLiUQZkcreie2sPoCNfQu6lwM
e6+pCPd+hdh+tp9vI2bZ113kK5BZcvG0eqsXQJhM01i9qI6tY2u9TJxQFw8ek+lD
p/l69jZf/zlqsqNiNOSO5d8E947IRslk9vdnBQIDAQABAoIBAHM1oT6hXxsO+haM
YYYD82pC6z9rTWUjTQTrxPl26tlS3OuGKm2e9OAbedD/jxh2FnV9G6m5HNVwwXc8
ZPsuXvXiyiafXVGUFznRRXnL18EIsBkGn1e5wTMcQ2/oWCcwCWIfUaVcMqJIMQL/
N/rWlMBkot4jRWIYgNMNfadz70OgZ5eSXbU1KjG3v/3pee/b/rhwQREPY+b8J1bZ
Mn6a8G/flH/0bah1rMoZF3Y3LbQteWoqcNECAReyQdAXFWoYCb7MBh9WEWtZMlB9
Qay17TqqK2fCc5sAgwRF0aGy0HrVfetwlNMQovCTy31coXbnC3mjk5j9i1FTsL5y
EVsNhS0CgYEA1PN7HYksVxODi6qb0xpFy+bozSfwiRY43z+Chx45hlocRokxeD13
Oe27Km8LrQFD961cP0jMMT2L9JSWxBmul0/3PR2rMz1hCkd3Feu59P2b5QaSA3ir
tQjeL07TUykDcGjM0efnFfyHTrmzfptVhK08DPR/NZjov9IA60nGvK8CgYEA0hGB
UgJDFQHW+lJkQnqcHi2dFtJn6kLxKUQLl9ggsNVJZ4UPMaIf0ZqbwiiX3MjK8ES+
YBM9aoBaJxxOYZBapChcpE6OBMxBoWH6EgNVuT41UtU1PQf7rQ1a7cFHjl8MJNin
ej4RzUvIrHA1tvwf5Lfro4zfdKhT9rqMJrllTIsCgYBBEVioELb8sFi8f1f3rApx
oE+4CCEmMiUUifpfQOwA0l3ba6YzAE0C5VIQSDgAF+flsRIDwEGsNgsio5hhuRsw
3t1DzmH2WSXily2bBBkg87EDzPsmlmOZAGmreOF3goureEpFRR/GBxBtDK7824fI
74nux3JrNRrBKfDf8/4GvwKBgQCsG4d2Z0AKsA0v3d0i0k2iNADmN9DHj70B8CwH
Zh2yg6y2Ub/XDFtBYYiHlLaHP0N1gvzuvhStcydpr7lQGfMJV8A/JdUZfTewPxOh
OnY7ZQpPKTuLG0VJzQ1YiehUbu8GKTNd5gizyIlLLkvz9bEztLQstThG3bqdBM3a
1D54bQKBgQC+osXwez9Bkee4egiDo8DL9gN6sB8XJPIjykKKv8H/HwPkfhjM/lt2
2dkThElDVEA36DWo6peERyDWIdetbwmzyLvD2KCAmXuQeNyYkjyRqmz6o1/LlajT
i1QtnI9+8pZKSAh3YrhmGui43RqJPSJv8NKJabQR2w0X1d1r36cFKA==
-----END RSA PRIVATE KEY-----`
	CaCertA = `-----BEGIN CERTIFICATE-----
MIIFgTCCA2mgAwIBAgIDEAISMA0GCSqGSIb3DQEBCwUAMFoxCzAJBgNVBAYTAlVT
MQ8wDQYDVQQIDAZEZW5pYWwxFDASBgNVBAcMC1NwcmluZ2ZpZWxkMQwwCgYDVQQK
DANEaXMxFjAUBgNVBAMMDSouZXhhbXBsZS5jb20wHhcNMTkwNTI5MjMxMzE5WhcN
MzkwNTI0MjMxMzE5WjBEMQswCQYDVQQGEwJVUzEPMA0GA1UECAwGRGVuaWFsMQww
CgYDVQQKDANEaXMxFjAUBgNVBAMMDSouZXhhbXBsZS5jb20wggIiMA0GCSqGSIb3
DQEBAQUAA4ICDwAwggIKAoICAQDH9KuhyhG+uhwXda576gUnklaAmQNDmAA/NBP+
qbx3MQIwDcfthikVaEvHPs1i6Roh8x35svenyq6ig2JE6Db77mq6/AvyrEuB2itw
aeB20KqVxm2jKc8gQMiJOoYzaTElceCLXSwM2zxxefiXkk7UcyH8o3tWHyQHuGqB
nqQa2FD/mlP6Jf4vxVnYVjab4k8ytFxqvWQGIqr1+qw3lY0dpNFetvSX0JUYzrtB
2RB5l1n6GeFfBBgzi8cLmBNLdHP8fpjb6J3uo2vGzMFKGAIP4MdO56XEPTYqXTxm
0Iw5uAOSf8ZxFhSToUZoOCPE3EKQh7EUyKn+bJSvDiIcXKwTMBMqoaLefx0yTfPS
hkv1wWMVZ4uLGvJ2epgSLMk3vTKycDMa1qVAL5+Kl4+W39LU/mq3I5NSATZVvwrk
SMGouLAMfc14hMfi/qw4vIhvxrulnkCkNdoLY75suUhaOgcvCeLD7XaBgeglfmCk
Mn7ewI+dPkEgQqbqSy72nul5URubNO5e+fQ5LlwoOUX22iA0YJ/98KrutWMIq1GB
+ZM23ZPzgL1OQOgPBzf3GacyWoQMZl57apJPVUeMoakuqdiVefR0w3LxsMXusXIs
FgtNIpj7/i7vc+CSnd1yp2cUFe3Lagg9jflSNEanR/rQoTxyOnjl4Q4zxtrhTIND
ylqDcQIDAQABo2YwZDAdBgNVHQ4EFgQU44b/zezM4Uz3S6WXtXN+cQoj46UwHwYD
VR0jBBgwFoAUVQpCveVxgNI3RhgPgillGF75pwkwEgYDVR0TAQH/BAgwBgEB/wIB
ADAOBgNVHQ8BAf8EBAMCAYYwDQYJKoZIhvcNAQELBQADggIBAAO1OOyIMjcFqBWB
svy/2lg87mGRhVnOlx8eeDw47Xrq2H665EgTxKefwpCe5w2vWvijiPenag6BGqbH
M7bwVoMf5i9ETtUihsFybST7QRQV8uaDk/uL5mLgZ5qucM+Uk1P3Y7Y4JiRxvLLn
0FRKTWfvQnql3KLdeHQ0nYKoiHIlLfrcVI9hDxHEs75ZJ/4OHMzyUDK+idQu2UXW
nY6Mydh1oMQ07vfbuzix7OntJB5/aP+XO5cTfqhUE68sXFHOuJrEidHO+Ec3nLib
kEN1hOhN2z3PGOisyg4GVrbz24y5NxEhE8qfFFjBDdHHLw42EjuZtVBIh0Ubh5d9
jxN1QUW0MdH0B3prb4D+ptab9LxZLe0prfv/iDqfVgtDZUrZG2xaqe9/aX85SMLn
4nx42zDBnNkgRaVAc5oC/WT4IhbJG2YlWZ4ZymvoLl4ZrKms983qcg6cK++MZYz/
ggC4eqSXqApQ99FA/j/c9Is3iMuWArhpVqe+sNjcD3+Fudhra91GxYAr6R+HubJV
hFI0ryZ+QphOSQMPi+ai5SWidwu2k7nzCwIQ7IkBBlwnjpVsZYtPXOdb4vf36ojU
rc8cdnePDG810YTvyE+nullZVnvQ4biHckIueMLdcxvg0H7BwI4rUFHxwckyRrHO
0kR1qoe1f8jDvCzcFeImPQ2BR4FW
-----END CERTIFICATE-----
-----BEGIN CERTIFICATE-----
MIIFpTCCA42gAwIBAgIUJJbxHzxLuXXB7kWNgnRK8jXRsEAwDQYJKoZIhvcNAQEL
BQAwWjELMAkGA1UEBhMCVVMxDzANBgNVBAgMBkRlbmlhbDEUMBIGA1UEBwwLU3By
aW5nZmllbGQxDDAKBgNVBAoMA0RpczEWMBQGA1UEAwwNKi5leGFtcGxlLmNvbTAe
Fw0xOTA1MjkyMzEzMThaFw0zOTA1MjQyMzEzMThaMFoxCzAJBgNVBAYTAlVTMQ8w
DQYDVQQIDAZEZW5pYWwxFDASBgNVBAcMC1NwcmluZ2ZpZWxkMQwwCgYDVQQKDANE
aXMxFjAUBgNVBAMMDSouZXhhbXBsZS5jb20wggIiMA0GCSqGSIb3DQEBAQUAA4IC
DwAwggIKAoICAQCgt1egggv5E5Me5UXDYgWR24Yh9foArVqKeoQWB251ldK1xAkg
Id6bBmhWFRBADDOjim7HNyc1abUF3R7JSdlpOXCTJ45+W1CR8ulviMt9SB05K4aK
RMU0OoFZrCps7jbv9+OIkhffefUk5kOx9l2bhPXhMm70wv1uV0BiT7FblBrWWdtd
98oX1xjCrWfC+vyJ7uVN7jxluxPZGSX6HerNLD5vR7J3vBnS6sdTbB35kaG+9VkQ
LCBkUy+vTBNepuTxc2Fzhv0UPjzGeGxwMqDgHSZ2pp1LVmey7xzV6JQHvbj/UgMr
HrvehS/xWZYKldQAsAjoMWHx1aMAyRkfq475zjKzsyUxh9ziUZTtYavuc+Z+vGnL
0SmFyoBYrlaVU20gkLi+sJtviWtt99/Z6+vWMab14BT6aUdAY9IDH/RtXgDgEXPs
CJrpSGavJA/lr9AP4fYnfOFywuwGt4R+J28FDoIZt62l7nv2eoHVCMQcEXGVI9yO
LWsBm1Ivp3vjGwQbI/V8lpaE5Jej91avLrkgnl+VOLrpf/uQnMcdshILtzfjMLwu
y1gGJjtjBrkeWnK9m+XcVTBRDTNcQ0depjfQqDWXPB5eHmVFV+sRTkiIsfmduA9Y
OLT5bzSY08kyrX6RT9DgKfm0ejlplJ1OyUeXZ93fK56qTECRjFvctU4LAwIDAQAB
o2MwYTAdBgNVHQ4EFgQUVQpCveVxgNI3RhgPgillGF75pwkwHwYDVR0jBBgwFoAU
VQpCveVxgNI3RhgPgillGF75pwkwDwYDVR0TAQH/BAUwAwEB/zAOBgNVHQ8BAf8E
BAMCAYYwDQYJKoZIhvcNAQELBQADggIBAI8vGfKSkuEri5JzjvG60CE0MnSU2+4T
cVQ0mB/XG0GZpnzP/t+XzTH/z0IbQDU1TNGLa3L7mBEYeFF4+1x0gKhJDc743qT0
I2GEJI6yhV+ftJmDcImfP8J9AVYxUYeKi/+WIHU+372xCLjSh8D4gvPETmRkDQuh
uQYZDXIEPaKHUMiRcgFMNNFkKdIILhXnPfx/sAbs7csJcldek91GEx64QSlKNTHN
e7m/9SXVx5JTAQrkk187j5I4GIcJAoj9yTzpXnqIhdl5hatJZmWb5LIr2ES3GUBf
61pJuTYmW0b4e5Qpx/5qGdGNvNazP9HniuJf6gBVquFnjy1Z5FJ196EAh4MAr4FB
RL9LIBTmVRWol7cS46K1C41+fIBNZE64kCOQs8tXjtqN3Ow+VZACw8WYSMmiHMCX
O+sG3U4K+I921rcgf5crz5TVUiCEe8rBiexGZ3q2EQAUSZM2bl9XJQD9IgDXUbq3
weQI0XQ65MCbxHA9MlugB0ZMdmpCXC4WlIpbhV7Ovu50W9ZrKaGzs+hYBSj1x3e8
Tzv3AsrIQ4Q847530qvpey+B4Yw2464aSfMr9hFT5OVfXwdXDF9DvRWIqr7Hjscj
GAdoQ+vDSkHtMptCtGoyHGcAVaKyw+b7KYtzkhsblSndiLqaUrPqfeNvvS5m/uHx
2EdJZi8UBOmQ
-----END CERTIFICATE-----`
	CaPrivateKeyA = `-----BEGIN RSA PRIVATE KEY-----
Proc-Type: 4,ENCRYPTED
DEK-Info: AES-256-CBC,5CCF7433030A5740972B840BF0C5A1E2

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
-----END RSA PRIVATE KEY-----`
	// Server certificate, private key and CA certificate
	TLSServerCertB = `-----BEGIN CERTIFICATE-----
MIIFPzCCAyegAwIBAgIDEAISMA0GCSqGSIb3DQEBCwUAMEQxCzAJBgNVBAYTAlVT
MQ8wDQYDVQQIDAZEZW5pYWwxDDAKBgNVBAoMA0RpczEWMBQGA1UEAwwNKi5leGFt
cGxlLmNvbTAeFw0xOTA1MjkyMzE0NDJaFw0yOTA1MjYyMzE0NDJaMFoxCzAJBgNV
BAYTAlVTMQ8wDQYDVQQIDAZEZW5pYWwxFDASBgNVBAcMC1NwcmluZ2ZpZWxkMQww
CgYDVQQKDANEaXMxFjAUBgNVBAMMDSouZXhhbXBsZS5jb20wggEiMA0GCSqGSIb3
DQEBAQUAA4IBDwAwggEKAoIBAQC2VeRleiVMqw/1zcMZrWJmML0nTdYaNliI7WIj
wQvZ+k05D469nClWg7vZ4whOK9lxIHzjWyxSC5FUFqX+NrKm3XbFo9/9u33GCjdO
2gaTr/zCter0a0kPsL5n9KwUkXseRMcrQyV39vl6c9Oecnr1LBhIvg1Z9iwFZsgS
aE8/NMZbsFALQx6EdiltG8lmz4jNPMAd2XS0MnKyO4EGXkkOvkthhf5kJ4KoGtVa
fhZUyActPeUYnaEspXxHE3ANerZPJs+Waykzt9oXpgplyQAmEPTzX3EttCLs0qH7
ZpdceYsOxSm06ea1G46yMzEoQMH5OLkmjfeohngDITa8ELy9AgMBAAGjggEiMIIB
HjAJBgNVHRMEAjAAMBEGCWCGSAGG+EIBAQQEAwIGQDAzBglghkgBhvhCAQ0EJhYk
T3BlblNTTCBHZW5lcmF0ZWQgU2VydmVyIENlcnRpZmljYXRlMB0GA1UdDgQWBBQ+
DNEDMyYWywiYdD8or9ScBq2JdDCBhAYDVR0jBH0we4AURfQW0owHXQLtnkMvKDSx
CY4Wy5qhXqRcMFoxCzAJBgNVBAYTAlVTMQ8wDQYDVQQIDAZEZW5pYWwxFDASBgNV
BAcMC1NwcmluZ2ZpZWxkMQwwCgYDVQQKDANEaXMxFjAUBgNVBAMMDSouZXhhbXBs
ZS5jb22CAxACEjAOBgNVHQ8BAf8EBAMCBaAwEwYDVR0lBAwwCgYIKwYBBQUHAwEw
DQYJKoZIhvcNAQELBQADggIBAAv98L5UBqLla4cpLQJQnW/iz9U0biLYRc80JLJY
50kZ2gsQeGxAd3XLzlgjh95z64JCaFO2XRcpNlEvJ0knAQby8UzBEvDYUDJXlG0p
4WDXndZDqXq3xMs3HHADT42KOLXKu1xcPZ/JyInLdZXnB4lMDIiPk2Kb/w6J+aIV
/jMMosxoEiOr1Rzv34xRNys/HpxFsItwJwUmMzXyM1eLDQPnvklt52va3WdZw6Ki
N1+KCxrbKCxBhTB1itvOAHX3KC7UTt0k0fP7pDc6lranc3hmDRgGIBy5vW+hrtZy
BGMJ/9jT+LWaBBAhhWQOYT0ILC/FmgUipC6kzB6KJrcDsrAxyr+lCxB0qy2HBtnm
s/geO3f6Z7cXLirwNi2SY6+D2jNbguJzvX43x/UNa4qUvXRi5tj/TwDth9M6jEnH
6ml+13liI8QX4+79ODgePt6VjxR71+KSd+51Qcwn1oajr80KoI0NhuuP4ZCBV53K
5Njz6XvJsPMgeVAIRmt3P+gOCW8kQrX6Qgb1fIvSrqR7vLpmybRQcvPt+k/T3m1F
VTkSsST68aOXCmHFLNA/6QB/OUatQYSr+O5iGz99xPEyz4EvHN5ZqTf51d+kyUje
zXsMLXBLCRwYz2U8y1FCcdOvCqDixneworIrBYDmoSafGUceSGQZOCN6ajW7fujg
FWy1
-----END CERTIFICATE-----`
	TLSServerKeyB = `-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQEAtlXkZXolTKsP9c3DGa1iZjC9J03WGjZYiO1iI8EL2fpNOQ+O
vZwpVoO72eMITivZcSB841ssUguRVBal/jaypt12xaPf/bt9xgo3TtoGk6/8wrXq
9GtJD7C+Z/SsFJF7HkTHK0Mld/b5enPTnnJ69SwYSL4NWfYsBWbIEmhPPzTGW7BQ
C0MehHYpbRvJZs+IzTzAHdl0tDJysjuBBl5JDr5LYYX+ZCeCqBrVWn4WVMgHLT3l
GJ2hLKV8RxNwDXq2TybPlmspM7faF6YKZckAJhD0819xLbQi7NKh+2aXXHmLDsUp
tOnmtRuOsjMxKEDB+Ti5Jo33qIZ4AyE2vBC8vQIDAQABAoIBAFFah3qDgkrv9EIP
GaLGSqYfzvXPc9zkLKKDkAs14gzS8kuoVl7hY0xXoKQ5+QWv6Ofhv/dNQbwlryUe
qdeAHjv3ijvqv1Edq6OGAWFsRAz3M8bIllmR2NpoWLxXXcQbmRxLQbZa0kfxaSSq
s1v55kSn9PnX2msuPDv4vQIBioI/T15Kj8pkPTD+xCM78MVkejHLykScYzP+sOra
6oz/+/CygZo8RL9b5VOyzFaBNDMgRjZImjA9WZgWcxvrYT9G13Th+h9ehVyZBzXm
zjf843/piLx53I7286DL3fwDnLcFtTam5qBJp24mZ24fUrIU26UeBD6S8VLpFh0m
EiGiXn0CgYEA6tXa3UbMscHS+uA6+WdMdhw1pjllSXk+Xevo1k099YdR57oFhnPk
VbLQADNn47+CbtP1VWT4K0ifbiog62DyrWD1Nr+twukp27BelTPtwRsNopyf6JdH
yaM2vHKK8QVeRTuonRaIdQTYDbcgyjZCvYNLNS4p+xmjt+PfV1cYsksCgYEAxsTB
X22s+tckQM3bGHV7LgWORnE1TVO2iRPtXHbmaeIDtIQkZyHsROf731uwkviMJ/VH
VN7j9nOaPNfXV4CFI63KFq+9fU5/6uyvYqUjnAMmnY0wcb0vZrX9S8iG05sefhJA
F2qF+MKAXv+nHz1blStiJOn2Al9lXROe2vDXKBcCgYAtEByPmaZt9eOmgKDQeyIR
CBjDbQAiuUFehIaJmQjqtKzi6q9kvZhzWXzQjQzCh11gJwxM9rMomyb83Ni0UkY/
PvvPJbkSTevaXF3KA8z74VUcfzGwUFdWOhumzdQbrAwK/Qe+HTTSP7u15g5Ev2TP
OQKVkeY8aQ8hmnihKfVzQQKBgAJpQptAvldwoqzyklTSoALn6FjsaKQeftlc14VG
n/bq2fann7VskOwpEEIeX574yuJ1ND9QUN2de2J/j8os8fCD3C64RQkvGZgk9N+y
dyMeWqqQmuTWUaviS0dPEGuitxT6bbKbBTMGUP5WBZX4bZL6qr8d43EV7Sna4hNv
1nxTAoGAGSL3JJsH8/OJAG2U1Uasq8yMEn7cNlH9FEalRFe0rDV6wfqiZ4R8vCcK
2HVOb8AnDf6LWQUjZZUlYDxM4luy3r0BmkmUqxHFdh+1wFkxnAi6GMB6VXPNkOAA
6Qyz5x0XKvcEaS2K8qjnPtZwh91EwO3p4QTtNoeXji1mJyggjoo=
-----END RSA PRIVATE KEY-----`
	TLSClientCertB = `-----BEGIN CERTIFICATE-----
MIIE4TCCAsmgAwIBAgIDEAITMA0GCSqGSIb3DQEBCwUAMEQxCzAJBgNVBAYTAlVT
MQ8wDQYDVQQIDAZEZW5pYWwxDDAKBgNVBAoMA0RpczEWMBQGA1UEAwwNKi5leGFt
cGxlLmNvbTAeFw0xOTA1MjkyMzE0NDNaFw0yOTA1MjYyMzE0NDNaMFoxCzAJBgNV
BAYTAlVTMQ8wDQYDVQQIDAZEZW5pYWwxFDASBgNVBAcMC1NwcmluZ2ZpZWxkMQww
CgYDVQQKDANEaXMxFjAUBgNVBAMMDSouZXhhbXBsZS5jb20wggEiMA0GCSqGSIb3
DQEBAQUAA4IBDwAwggEKAoIBAQDNRc9LszNAe0A2nZMXCA8zhl0kf1BGdozORrMG
Hbs38edMCr5XrsNteFIhp/iDRRd/8xvfV6WpY1qtulf37j155wB8Uehp2lRk4WaD
ZAK+h9y+ekBW0thpKr38vPTBVUEU/50AsNetI8QbSMh/ZBv9n00u7YFBw3GY/pAp
Hr/1lSX4Ntx3bQM5RhKsSq02ft1wwrUe4reLRXPNFj5BUAeS+JtGUMoiAw54yQ/+
1VFRgGA1zsKEm/A5bPxkEpum/fiE2pu7OfYX53I0phh561ITVp6QCaipPH6OWi6Q
Tb/+zScK9wZc4JNIOghoe0QaGQdnDNcLKsgVs5r1ZCAaRcExAgMBAAGjgcUwgcIw
CQYDVR0TBAIwADARBglghkgBhvhCAQEEBAMCBaAwMwYJYIZIAYb4QgENBCYWJE9w
ZW5TU0wgR2VuZXJhdGVkIENsaWVudCBDZXJ0aWZpY2F0ZTAdBgNVHQ4EFgQU5sVj
9w4KkxrIDx1g2hm4V/LhBnwwHwYDVR0jBBgwFoAURfQW0owHXQLtnkMvKDSxCY4W
y5owDgYDVR0PAQH/BAQDAgXgMB0GA1UdJQQWMBQGCCsGAQUFBwMCBggrBgEFBQcD
BDANBgkqhkiG9w0BAQsFAAOCAgEAKRmPDANC/yBAbk7eQM6XEIUzgm4y0/rycXgf
Cf2fQc3SvvG92hZbKv9sS9nWd0Sst0S8Iq2ZiK9PqvDzobVJ5HTHY3Oo2X5n3xmj
lGBlSI9oWaOj8aeG+8gkcJ3ruG3IHAlA70yws2pMbD4VxXf0BaWHf/jOuqrF4ykX
Fh67VxPL3NQkAdrkuocjw8/azUrBlLGnsl3eih9Me4IlsbxkHbPMFt0RL98wodVI
MvcKobxSoB7ak6OfMwV14FM285xbEFcLgmrXPnfQwieRRbTPXeeu/AecX58tVVRc
Edh8I0w+34m2Gqstl9PNzccnlSSUzef58nuilIH5i0XMoLIlpMhvOMCrUf7zkrgg
tt8kiFw71BoBm5mV6yxXjjV8j0lSdFiRSYWcmMHJzmR7sL58XS1w35k/dN8hUedE
QXGeQT1UUjm0XiktrqlAOQGBp4MBHazxLE5Qywozb7QoIeq1/4qLbZfuPXjg7aIC
C4CM6wQxYRD2P0HO4FLM8goxpVG3ShzbFSgPijcyAXSD/I8k9RN5hUsDN/yXpFuD
hdkUxbG4zkSzu073qxebiDJ9IBO3qP3yFURWlYJPZxSTb2e9QDHExKaerKRQ7pAw
DsKqT6LVcsxVzwb+TFP3z7meOX/iSxJPC5JRn6fwdPDCUylCG522JedkoLz3t2V/
nx5hGa0=
-----END CERTIFICATE-----`
	TLSClientKeyB = `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAzUXPS7MzQHtANp2TFwgPM4ZdJH9QRnaMzkazBh27N/HnTAq+
V67DbXhSIaf4g0UXf/Mb31elqWNarbpX9+49eecAfFHoadpUZOFmg2QCvofcvnpA
VtLYaSq9/Lz0wVVBFP+dALDXrSPEG0jIf2Qb/Z9NLu2BQcNxmP6QKR6/9ZUl+Dbc
d20DOUYSrEqtNn7dcMK1HuK3i0VzzRY+QVAHkvibRlDKIgMOeMkP/tVRUYBgNc7C
hJvwOWz8ZBKbpv34hNqbuzn2F+dyNKYYeetSE1aekAmoqTx+jloukE2//s0nCvcG
XOCTSDoIaHtEGhkHZwzXCyrIFbOa9WQgGkXBMQIDAQABAoIBAG3WuQGMthP+33Fa
B+b/DQNJFnX9GftaXCXUdt8C8bcR3e42oXtRrIjbJTeVJck4I0b30yZDRAXLgC4N
1Bx6grLEBOKBAZgCl74TTkNoNH+3O0tBJ0RONjawBFFE2sLI17ZgwKYp+n5O8RSY
cZCZBNFFeItv0wZZimOPc5xNM/I4ErRLwGt2rxoGv1rKboWmdVmgpou1lu9NOk4i
35p69dRVJBHMHSos+irdUaOMDYehJfq2POv0fBaTdfki9Hev67+BVsG0Sv/7wc8z
cEkQiYKzREqp3n3n3kX+Y9z1shEVGiAIhmvxF5d8KadUc/ZYt0US0LZ9y/RT23IS
60KF3DECgYEA97lB+IvV34jcqcVx1SLh45djqzDjoMt+R1G1UNgePw4EK81jyybM
S736SWW47NcpQm1mBC/yR86E/8HTvJeDfG5RzNeQ3EL/YD54dMby82Z0DiWhBh28
MDDafA3Mz01vsa6JccNQwRSxJ7LHXnHmFkSiq7uXj4wjcE5szCriDnMCgYEA1CF5
t77rmPPo3AcZt8BMGR9DLiotcJegU1K5GZTNGRakqC/z3Z4i9dTjLuVITVMEWcLy
dI3xlANsWSpfAl+o8lqXtosLU3EqwDVFcQhgGZ82qy2n9lKShp7unblNLMvTo2bc
yXL92gXWe/Ffe99pefvTZPjcWZsdLVMV1SLrhMsCgYA5WzWBHK9qUD/3NvYCSU9P
6M905Z7ure8RCEQY8dEe/FnO4oFGmjcXGmeG8vx8Kd0tujKyqX9JTPHSXchulA4n
k0txSEAMH58NY6l0MFQ0MvaQB/pedigKaGVN78wJ+33u/+Bm7LgX3HigHm622VxS
r5WSeZ6/58yUxjO+mfc5LQKBgEDb82FRCKv9d6c+rIhYTWwo+Nt2neodEjInytyd
eJBBMhsSfle2cC4F52iBRjgON/hR+NvWQpTk7w3cPKx3HyrqtBRmMxJSOaYHI7JJ
w0hSO97e2MohuRdcJM0oyQX3VEBTxRH7DM7KlifTR0SSrKdVbe8jgAwNt0ASdUxz
nts9AoGBANNxMkdxeDN1fHUDq/lvtCwQw9s/XukDHJgOrQC6Z4fjBSTZFIb/MEgX
vkYHo46mecOWXZZdwDgjUmDQCNdslWvooRsrpSp4Xda3zRkGbzzk6cL0AESgFgBO
Gow8ytlGjhTpn1hL+f+2cv44sqdZY6hThuBpBcj0EmM6seWu2GaF
-----END RSA PRIVATE KEY-----`
	CaCertB = `-----BEGIN CERTIFICATE-----
MIIFgTCCA2mgAwIBAgIDEAISMA0GCSqGSIb3DQEBCwUAMFoxCzAJBgNVBAYTAlVT
MQ8wDQYDVQQIDAZEZW5pYWwxFDASBgNVBAcMC1NwcmluZ2ZpZWxkMQwwCgYDVQQK
DANEaXMxFjAUBgNVBAMMDSouZXhhbXBsZS5jb20wHhcNMTkwNTI5MjMxNDM5WhcN
MzkwNTI0MjMxNDM5WjBEMQswCQYDVQQGEwJVUzEPMA0GA1UECAwGRGVuaWFsMQww
CgYDVQQKDANEaXMxFjAUBgNVBAMMDSouZXhhbXBsZS5jb20wggIiMA0GCSqGSIb3
DQEBAQUAA4ICDwAwggIKAoICAQDIFHI/+8ZhJJPuktwXLG9ejV8R3MRS7+0KvJut
SEZASZbE+x4/bvrxy7Y3ZmoY8W7r/HxAqj9xmvfmqo355Ix4PqK2S/v2jiVpLf5o
DTEtOI63SSLUvrOdmhjnQz+ZMqLZCWCIQcCXSCfhBWWRyH+iOpAcaLOOaoZI/qCR
FCqOCmxh3k9Kthex7gwTNCqaZ+jrhq603HrfZ7DtRolAX7oz4ucDC9qraVvGx6MO
nNh5E0qvOHhuZ1sExWyG9NqEkPaqfDeK0svsiHPWdm6jvo0r/8DiDKeda+MvM1C/
ZQxB+paH0qmTqxnsA3AngP1w3NzGvHxCXFwnV3D5iMTayohpc0dyZuaGUMKWKj5R
jq0LwteB1uAyyhEz2SG8qJL+QbawM1qxK1EjJOFvk/pdTnN7HzDc/+bklUnYBbbd
vdaF+JAdYpDEZlMrHTVfeY7kMM9TbHQbmsmNoUr/GHzDCqoiEa22DITJwSRf2rUe
VISsvGCyXBSyUilK5dUWmFl7AXprpMBIKQL9X50Ssfx3rs2mWGEP0YySFRn0zrxl
wg1yBygh8HC+HGNaMdXZDGXeSJzITQI1anngHhkOaOh29UEGM2k5SNiWOZxQEV4J
L5BBjaJBCifAZcCZd3sPkxmVurZi9vAVc9JT+Avkg4UuDu/HWYpeE302S3l5QSKI
eVeNEwIDAQABo2YwZDAdBgNVHQ4EFgQURfQW0owHXQLtnkMvKDSxCY4Wy5owHwYD
VR0jBBgwFoAUd3dG9VBzLMLYT6z3+LT/U3p1SpswEgYDVR0TAQH/BAgwBgEB/wIB
ADAOBgNVHQ8BAf8EBAMCAYYwDQYJKoZIhvcNAQELBQADggIBAIUilHtj2m5qV+xc
ElVY3Gsv0mhxHNbEkG+EgOoWd1QWjvz58gkddgGxAb5mJYM5FSdw4yoSIuUYy6F0
dXxAElWFplNiPY46OJ2/MwINGNZWfkI7jxVCYanHXGJa5llVpPNhLpJEVGt8FOkd
nLu2ZKCLLSAOC9Z3R/FNxF4HiAuN8Z3OYX7jtaUSdgPogsjzsuWWxuy5Rs8A8hM4
483UfakoHjLtXAzbdQ1sA7k8YvY3u1t2b3x9jDEcHmz2FYX+N0BtzgSsZvLxrT9b
wUC+g4Lspl/Lnp8Jrg8k/DQdcC1g0rVMi5nLNtHbH/2j502Na/TDtPANfAlf9gcd
9TjLsxqax8L7vXvrrcvnJTDZZVjA3NQhDV5EFuwjcSoUq02p7c63FYChNNCdEjKs
8MA+jH97xafhR6TvxW9R9BTcIwrrJmgmQ+b2hz9sqlE1ZDT4Biy4bUQiAvRTuSxX
ch+jNPLD3kyfDEixhE1+5luutC77b98qc3KWsG6l3XDQQf/YZ6h6iIkZsXorGtxg
sLcSOZBc3XyP5twMeOw2ZOMC0qLupFL2MBEmKerlHo5ehQpW16KBHWn1HxFL8j24
PAsalRNQlxxWYCEYsf60TIUSqtyt1P5G7S40Rn3CP9SnoX6Q3E0POxEGFe3SStAY
oCvHkuhGyVKRT4Ddff4gfbvMPlls
-----END CERTIFICATE-----
-----BEGIN CERTIFICATE-----
MIIFpTCCA42gAwIBAgIUJuIWrqbNSCJFLYQi0C72k1lG/r8wDQYJKoZIhvcNAQEL
BQAwWjELMAkGA1UEBhMCVVMxDzANBgNVBAgMBkRlbmlhbDEUMBIGA1UEBwwLU3By
aW5nZmllbGQxDDAKBgNVBAoMA0RpczEWMBQGA1UEAwwNKi5leGFtcGxlLmNvbTAe
Fw0xOTA1MjkyMzE0MzlaFw0zOTA1MjQyMzE0MzlaMFoxCzAJBgNVBAYTAlVTMQ8w
DQYDVQQIDAZEZW5pYWwxFDASBgNVBAcMC1NwcmluZ2ZpZWxkMQwwCgYDVQQKDANE
aXMxFjAUBgNVBAMMDSouZXhhbXBsZS5jb20wggIiMA0GCSqGSIb3DQEBAQUAA4IC
DwAwggIKAoICAQC+jF19L7KSozBUCrkrsjZsgLMzHHnNBcwV3jBKeDHUJxnW7KfZ
szZ6zeC6aWd0nMUv/xCCUyGaj21nuXh8uHu+oMS+JwpAOqdjlhC4NQJdgxudgs/R
f1YNIUTZ3X4Ka8vxnM3XkCK4UHIlCaWYOfOA2d1MfucdJSAAbDL9v12ZYkT8OTAM
7PsozvbEVBpKM2YMinTXMMNIdyIWhRvwy1Hqe4YA/GiU03ouJFmE66Rha1bIa+lI
KuGAevBO+uk80pDiJ/PLyrR8J08leVqb7kSz8Bc6vk319gFoWwIHgW04dl1OdS4u
l4vJYfpZcSyF+RY2Actdy2nd4iBqtuK9FCrI3Vkqq5OSqKw0zskJ20MFQoAi8g0z
CJCCHMPod0a3ydghk+bUxHnEFnZilgXtVlN4BAkFLEUmhGYDyb6nwB+vqIjmhdj9
+i9d2IXmUME2iQxlMwnZdK9hCGVfwCTuVqZ4soqYCZ4cixh65PPstHTAGB/ru+JY
VztS8nhJLygXkKAXfrapcfbpmf7BZFvPaR1fbn+eCHNp9khY597MwMM5QQ07mskF
EPQlzURodFaGGaRF+QA6BMfCFMF8JgCLRzaYJn0qQ1SvkbeMDlzaw2/fh9TwlFfx
okhnY9YzVkG0TwuT3yUvaV8DxJ6bdDI75jqN7hGAubA/8WcAEMyj7C72swIDAQAB
o2MwYTAdBgNVHQ4EFgQUd3dG9VBzLMLYT6z3+LT/U3p1SpswHwYDVR0jBBgwFoAU
d3dG9VBzLMLYT6z3+LT/U3p1SpswDwYDVR0TAQH/BAUwAwEB/zAOBgNVHQ8BAf8E
BAMCAYYwDQYJKoZIhvcNAQELBQADggIBABpWo4tpHTJqtwnHD7lmqhIZF+lB6S8b
wp2XwtqzaxsSWk9yIPqySOLauOKkA3r1RMWlu05qhSrkI5AHbNYHhAOtKiY7IGkw
p9PE5m52uTIlJrFuUE37tmvao8+hRHLSCxVB++gOfbKAlvAsZIBNoCosYsHcwWQH
1IAdKpQVmSs4Bf5ZpGx0HDhuzLp2+GMLyzo3ZZDKdG4euTEQUC2+w16w4f2XPkob
XTPT/hSBld2SLEln5rYIIcSu3sBgMgE4keG86H2TQzHsUX+Md1eL2BWmxR7MWEyz
hA8poTPyj1f99eZVgiiu6BBhrjzoCUX9w4XR5XDgibEkbfzlsZzTi3tvUIh3fz9m
q6/XckeOiGib34nDI3lUoIoLj9HIESAsuMyuRW49tAjax4IlVDmPttLGY1Yo0mhu
QJxUqu6OyaR6HZCDzIttTIE5jmAF/FjgqvHY3qn8jJyBbMqw7zMVYrJT6lGfGVLs
FyFwZ8+yj3lhceWx72N5k0cUE2ytSmDoH2xgbLicAHZPNI/Jjubi+j4d/HeO61sc
JLCwEcHpScz8aIqIDqHkLQKTbremSDrde8uW1KG950rdcz9ykcj6yOwrLEgsl4uF
mfnY0G770B8jHdU2wsnisdz/94Y8J/uw2FCtV3fdYPjDMeovG9vDh5niik/AbQDr
TcMxvnJJ8CST
-----END CERTIFICATE-----`
	CaPrivateKeyB = `-----BEGIN RSA PRIVATE KEY-----
Proc-Type: 4,ENCRYPTED
DEK-Info: AES-256-CBC,9D455985DCE5A63E21DCAC5AD644194A

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
-----END RSA PRIVATE KEY-----`
)
