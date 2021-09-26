package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/machinebox/graphql"
)

// func (data SpacexLauchesPast) spacexLauchesPast() SpacexLauchesPast {
// 	return data
// }

func handler(w http.ResponseWriter, r *http.Request) {
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
		fmt.Fprintf(w, "%d %s", i, key.LaunchSite.SiteNameLong)
	}
}
func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", handler).Methods("GET")
	return r
}
func main() {
	r := newRouter()
	log.Fatal(http.ListenAndServe(":8080", r))
}
