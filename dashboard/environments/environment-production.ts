const environment = {
  production: true,
  API_URL: process.env.NEXT_PUBLIC_API_URL
    ? process.env.NEXT_PUBLIC_API_URL
    : '',
  GA_TRACKING_ID: 'G-9HF3HT6S6W',
  SENTRY_URL:
    'https://b4b98ad60a89468284cf8aa5d66cf2cd@o1267000.ingest.sentry.io/4504797672701952'
};

export default environment;
