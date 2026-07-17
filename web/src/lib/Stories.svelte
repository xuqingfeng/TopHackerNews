<script>
  import { ApolloClient, InMemoryCache, HttpLink, gql } from "@apollo/client/core";
  import Footer from "./Footer.svelte";

  const API_URL = import.meta.env.VITE_API_URL;

  let stories = [];
  let currentPage = 1;
  let limit = 15;
  let offset = (currentPage - 1) * limit;

  const cache = new InMemoryCache({
    resultCacheMaxSize: 1000,
  });

  function fetchStories() {
    stories = [];
    offset = (currentPage - 1) * limit;
    const client = new ApolloClient({
      link: new HttpLink({
        uri: `${API_URL}/graphql`,
      }),
      cache: cache,
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

<div class="thn-stories">
  {#if stories.length == 0}
    <div class="alert alert-info">
      <span class="loading"></span>
      Loading...
    </div>
  {:else}
    <ul>
      {#each stories as story (story.id)}
        {#if story.type !== "job"}
          <li>
            <div>
              <p>
                {story.title}{#if story.url }: <a href={story.url} target="_blank">{story.url}</a>{/if}
              </p>
            </div>
            <div>
              <p class="thn-meta">
                {story.score} points |
                <a
                  href="https://news.ycombinator.com/item?id={story.id}"
                  target="_blank">Comments ↗</a
                >
                |
                <time datetime={story.time}>Time: {convertToDateTime(story.time)}</time>
              </p>
            </div>
          </li>
        {/if}
      {/each}
    </ul>
    <div class="thn-bottom-bar">
      <Footer />
      <div class="btn-group thn-pagination">
        {#if currentPage != 1}
          <button class="btn btn-primary btn-ghost" on:click={prev}>«</button>
        {/if}
        <button class="btn btn-default btn-ghost disabled">{currentPage}</button>
        <button class="btn btn-primary btn-ghost" on:click={next}>»</button>
      </div>
    </div>
  {/if}
</div>
