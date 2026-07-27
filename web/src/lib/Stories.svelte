<script>
  import { ApolloClient, InMemoryCache, HttpLink, gql } from "@apollo/client/core";
  import Footer from "./Footer.svelte";

  const API_URL = import.meta.env.VITE_API_URL;

  const client = new ApolloClient({
    link: new HttpLink({
      uri: `${API_URL}/graphql`,
    }),
    cache: new InMemoryCache({
      resultCacheMaxSize: 1000,
    }),
  });

  let stories = [];
  let currentPage = 1;
  let limit = 15;
  let loading = true;
  let error = null;
  let hasNextPage = true;

  function fetchStories() {
    loading = true;
    error = null;
    const offset = (currentPage - 1) * limit;

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
        const fetched = data.data?.topStories ?? [];

        if (fetched.length === 0 && currentPage > 1) {
          currentPage -= 1;
          hasNextPage = false;
          return;
        }

        stories = fetched;
        hasNextPage = fetched.length >= limit;
      })
      .catch((err) => {
        error = err?.message ?? "Failed to load stories";
      })
      .finally(() => {
        loading = false;
      });
  }

  fetchStories();

  function next() {
    if (!hasNextPage || loading) {
      return;
    }
    currentPage += 1;
    fetchStories();
  }

  function prev() {
    if (currentPage === 1 || loading) {
      return;
    }
    currentPage -= 1;
    hasNextPage = true;
    fetchStories();
  }

  function first() {
    if (currentPage === 1 || loading) {
      return;
    }
    currentPage = 1;
    hasNextPage = true;
    fetchStories();
  }

  function convertToDateTime(time) {
    return new Date(time * 1000).toLocaleString();
  }
</script>

<div class="thn-stories">
  {#if loading && stories.length === 0}
    <div class="alert alert-info">
      <span class="loading"></span>
      Loading...
    </div>
  {:else if error && stories.length === 0}
    <div class="alert alert-error">
      {error}
      <button class="btn btn-default" on:click={fetchStories}>Retry</button>
    </div>
  {:else}
    {#if loading}
      <div class="alert alert-info thn-loading-banner">
        <span class="loading"></span>
        Loading...
      </div>
    {/if}

    {#if error}
      <div class="alert alert-error">
        {error}
        <button class="btn btn-default" on:click={fetchStories}>Retry</button>
      </div>
    {/if}

    <ul>
      {#each stories as story (story.id)}
        {#if story.type !== "job"}
          <li>
            <div>
              <p>
                {story.title}{#if story.url }: <a href={story.url} target="_blank" rel="noopener noreferrer">{story.url}</a>{/if}
              </p>
            </div>
            <div>
              <p class="thn-meta">
                {story.score} points |
                <a
                  href="https://news.ycombinator.com/item?id={story.id}"
                  target="_blank"
                  rel="noopener noreferrer">comments ↗</a
                >
                |
                <time datetime={story.time}>created at: {convertToDateTime(story.time)}</time>
              </p>
            </div>
          </li>
        {/if}
      {/each}
    </ul>

    {#if !loading}
      <div class="thn-bottom-bar">
        <Footer />
        <div class="btn-group thn-pagination">
          {#if currentPage > 2}
            <button class="btn btn-primary btn-ghost" on:click={first} title="First page">««</button>
          {/if}
          {#if currentPage != 1}
            <button class="btn btn-primary btn-ghost" on:click={prev} title="Previous page">«</button>
          {/if}
          <button class="btn btn-default btn-ghost disabled">{currentPage}</button>
          {#if hasNextPage}
            <button class="btn btn-primary btn-ghost" on:click={next}>»</button>
          {/if}
        </div>
      </div>
    {/if}
  {/if}
</div>
