package main

import (
	"context"
	"fmt"

	"github.com/machinebox/graphql"
)

// func (data SpacexLauchesPast) spacexLauchesPast() SpacexLauchesPast {
// 	return data
// }

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
	ctx := context.Background()
	resp := Data{}
	if err := graphqlClient.Run(ctx, graphqlRequest, &resp); err != nil {
		panic(err)
	}

	for i, key := range resp.LaunchesPast {
		fmt.Println(i, key.LaunchSite.SiteNameLong)
	}
}
