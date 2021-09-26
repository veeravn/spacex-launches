package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/machinebox/graphql"
)

func handler(w http.ResponseWriter, r *http.Request) {
	//create graphql client
	graphqlClient := graphql.NewClient("https://api.spacex.land/graphql/")
	graphqlRequest := graphql.NewRequest(`
    {
  launchesPast(limit: 10) {
  mission_name
  launch_date_local
  launch_site {
  site_name_long
  }
  links {
  article_link
  video_link
  }
  rocket {
  rocket_name
  first_stage {
    cores {
      flight
      core {
        reuse_count
        status
      }
    }
  }
  second_stage {
    payloads {
      payload_type
      payload_mass_kg
      payload_mass_lbs
    }
  }
  }
  ships {
  name
  home_port
  image
  }
  }
  }
    `)
	ctx := context.Background()
	var resp = Data{}
	if err := graphqlClient.Run(ctx, graphqlRequest, &resp); err != nil {
		panic(err)
	}
	for i, key := range resp.LaunchesPast {
		if i == 0 {
			w.Write([]byte("["))
		}
		bytes := serializeJSON(key)
		w.Write(bytes)
		if i == len(resp.LaunchesPast)-1 {
			w.Write([]byte("]"))
		} else {
			w.Write([]byte(","))
		}
	}
}
func serializeJSON(key LaunchesPast) []byte {
	bytes, err := json.Marshal(LaunchesPastTable{
		SiteNameLong: key.LaunchSite.SiteNameLong,
		MissionName:  key.MissionName,
		RocketName:   key.Rocket.RocketName,
	})
	if err != nil {
		panic(err)
	}
	return bytes
}
func newRouter() *mux.Router {
	r := mux.NewRouter()

	// Declare the static file directory
	staticFileDirectory := http.Dir("./assets/")
	// Declare the handler, that routes requests to their respective filename.
	// The fileserver is wrapped in the `stripPrefix` method, because we want to
	// remove the "/assets/" prefix when looking for files.
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	// The "PathPrefix" method acts as a matcher, and matches all routes starting
	// with "/assets/", instead of the absolute route itself
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")
	r.HandleFunc("/launchesPast", handler).Methods("GET")
	return r
}
func main() {
	r := newRouter()
	log.Fatal(http.ListenAndServe(":8080", r))
}
