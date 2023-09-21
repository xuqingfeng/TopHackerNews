<script>
  import { ApolloClient, InMemoryCache, gql } from "@apollo/client/core";

  const API_URL = import.meta.env.VITE_API_URL;

  let stories = [];
  let currentPage = 1;
  let limit = 15;
  let offset = (currentPage - 1) * limit;

  function fetchStories() {
    stories = [];
    offset = (currentPage - 1) * limit;
    const client = new ApolloClient({
      uri: `${API_URL}/graphql`,
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

  function convertToDateTime(time) {
    return new Date(time * 1000).toLocaleString();
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
              {story.title} -
              <a href={story.url} target="_blank">{story.url}</a>
            </p>
          </div>
          <div>
            <p>
              By: <a href="https://news.ycombinator.com/user?id={story.by}" target="_blank">{story.by} ↗</a> |
              <a href="https://news.ycombinator.com/item?id={story.id}" target="_blank">comments ↗</a> |
              <!-- TODO: convert to datetime -->
              <time datetime={story.time}>Time: {story.time}</time>
            </p>
          </div>
        </li>
      {/each}
    </ul>
    <div class="grid -right">
      <div class="btn-group">
        {#if currentPage != 1}
          <button class="btn btn-primary btn-ghost" on:click={prev}>prev</button
          >
        {/if}
        <button class="btn btn-default btn-ghost disabled">{currentPage}</button
        >
        <button class="btn btn-primary btn-ghost" on:click={next}>next</button>
      </div>
    </div>
  {/if}
</div>
