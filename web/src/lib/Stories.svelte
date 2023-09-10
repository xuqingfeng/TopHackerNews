<script>
  import { ApolloClient, InMemoryCache, gql } from "@apollo/client/core";

  let stories = [];

  const client = new ApolloClient({
    uri: "http://localhost:8080/graphql",
    cache: new InMemoryCache(),
  });
  client
    .query({
      query: gql`
        query topStories {
          topStories(offset: 0, limit: 10) {
            id
            kids
            score
            time
            title
            type
            url
            by
          }
        }
      `,
    })
    .then((data) => {
      stories = data.data.topStories;
    });
</script>

<div>
  {#if stories.length == 0}
    <div class="alert alert-info">
      <span class="loading" />
      Loading...
    </div>
  {:else}
    <ul>
      {#each stories as story (story.id)}
        <li>
          <div>
            <p>
              {story.title} - <a class="" href={story.url}>{story.url}</a>
            </p>
          </div>
          <div>
            <p>
              By: {story.by}
              <time datetime="">Time: {story.time}</time>
            </p>
          </div>
        </li>
      {/each}
    </ul>
  {/if}
</div>
