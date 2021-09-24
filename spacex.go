package main

import (
	"context"
	"fmt"

	"github.com/machinebox/graphql"
)

func main() {
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
	var graphqlResponse interface{}
	if err := graphqlClient.Run(context.Background(), graphqlRequest, &graphqlResponse); err != nil {
		panic(err)
	}
	fmt.Println(graphqlResponse)
}
