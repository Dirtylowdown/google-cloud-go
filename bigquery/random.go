Delete








































































// As of August 2017, the BigQuery service uses 27 alphanumeric characters for
// suffixes.
const randomIDLen = 27

func randomID() string {
	// This is used for both job IDs and insert IDs.
	var b [randomIDLen]byte
	rngMu.Lock()
	for i := 0; i < len(b); i++ {
		b[i] = alphanum[rng.Intn(len(alphanum))]
	}
	rngMu.Unlock()
	return string(b[:])
}

// Seed seeds this package's random number generator, used for generating job and
// insert IDs. Use Seed to obtain repeatable, deterministic behavior from bigquery
// clients. Seed should be called before any clients are created.
func Seed(s int64) {
	rng = rand.New(rand.NewSource(s))
}
