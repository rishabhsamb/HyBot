# fss-bot

fss-bot is a multipurpose Discord bot built with response speed in mind. Written in Go, it primarily relies on Google's Cloud Firestore for storing necessary media.

Currently, fss-bot supports two main functionalities:
- **Outbursts**: Predetermined messages sent when prompted by certain keywords. Because of Discord's powerful image and video embeds, this enables users to reference videos, images, or other internet memes seamlessly in conversation.
- **Weather**: When prompted, the bot will query [a public graphql api](https://graphql-weather-api.herokuapp.com) for weather statistics in a provided city. HyBot uses [Khan/Genqlient](https://github.com/Khan/genqlient) for generating type-safe GraphQL queries at compile-time.


In the future, fss-bot is planning to support:
- **Twitter webhooks**: A notification service where users can register for notifications when certain public accounts tweet. Other Twitter integrations to come as well.
- **Stock & Crypto**: Sourcing daily stock updates for user-subscribed tickers.

## Cloning and Running fss-bot Yourself

You'll need a .env file specifying a few variables if you want to run HyBot as is. `DISCORD_TOKEN` is just the token you get with every Discord bot. `GOOGLE_PROJECT_ID` is what you think it is, and `GOOGLE_APPLICATION_CREDENTIALS` is the absolute path to a service-account.json that has permissions to run Cloud Firestore.

After that, `go run .` in the root directory will run the bot.
