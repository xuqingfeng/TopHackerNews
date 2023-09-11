<script>
  import { ApolloClient, InMemoryCache, gql } from "@apollo/client/core";

  let stories = [];
  let currentPage = 1;
  let limit = 10;
  let offset = (currentPage - 1) * limit;

  function fetchStories() {
    stories = [];
    const client = new ApolloClient({
      uri: "http://localhost:8080/graphql",
      cache: new InMemoryCache(),
    });
    client
      .query({
        query: gql`
        query topStories {
          topStories(offset: ${offset}, limit: ${limit}) {
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
  }

  fetchStories();

  function next() {
    currentPage += 1;
    fetchStories();
  }

  function prev() {
    currentPage -= 1;
    fetchStories();
  }
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
    <div class="grid -right">
      <div class="btn-group">
        {#if currentPage != 1}
          <button class="btn btn-primary btn-ghost" on:click={prev}>prev</button>
        {/if}
        <button class="btn btn-info btn-ghost">{currentPage}</button>
        <button class="btn btn-primary btn-ghost" on:click={next}>next</button>
      </div>
    </div>
  {/if}
</div>
