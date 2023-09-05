<script>
  let stories = [];
  const topStoriesQuery = `
    query topStories {
      topStories(offset: 0, limit: 10) {
        id
        kids
        score
        title
        type
        url
        by
      }
    }
  `;
  fetch("http://localhost:8080/query", {
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
  <ul>
    {#each stories as story (story.id)}
      <li>{story.title} - {story.url}</li>
    {/each}
  </ul>
</div>
