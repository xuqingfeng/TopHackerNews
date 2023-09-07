<script>
  let stories = [];
  const topStoriesQuery = `
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
  `;
  fetch("http://localhost:8080/graphql", {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Accept: "application/json",
    },
    body: JSON.stringify({ query: topStoriesQuery }),
  })
    .then((r) => r.json())
    .then((data) => (stories = data.data.topStories));
</script>

<div>
  <!-- {#if $stories.loading}
  <div></div>
  {:else if $stories.error}
  <div>Error!</div>
  {:else if $stories} -->
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
  <!-- {/if} -->
</div>
