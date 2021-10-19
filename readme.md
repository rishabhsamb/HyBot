# HyBot

HyBot is a multipurpose Discord bot built with response speed in mind. Written in Go, it primarily relies on Google's Cloud Firestore for storing necessary media.

Currently, HyBot supports two main functionalities:
- **Outbursts**: Predetermined messages sent when prompted by certain keywords. Because of Discord's powerful image and video embeds, this enables users to reference videos, images, or other internet memes seamlessly in conversation.
- **Weather**: When prompted, the bot will query [a public graphql api](https://graphql-weather-api.herokuapp.com) for weather statistics in a provided city. HyBot uses [Khan/Genqlient](https://github.com/Khan/genqlient) for generating type-safe GraphQL queries at compile-time.


In the future, HyBot is planning to support:
- **AniList**: The bot would use AniList's GraphQL API to source information on relevant shows that are referenced in Discord messages.
- **Twitter webhooks**: A notification service where users can register for notifications when certain public accounts tweet. Other Twitter integrations to come as well.
- **Stock & Crypto**: Sourcing daily stock updates for user-subscribed tickers.

## Cloning and Running HyBot Yourself

Coming soon...I really need a free evening.
