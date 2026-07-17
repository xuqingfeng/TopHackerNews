<script>
  const WEB_URL = import.meta.env.VITE_WEB_URL;
  const THEME_NAME = "theme";
  const THEMES = ["light", "dark", "dark-grey", "solarized-dark"];
  const THEME_BODY_CLASSES = {
    light: [],
    dark: ["dark"],
    "dark-grey": ["dark-grey"],
    "solarized-dark": ["solarized-dark"],
  };

  const storedTheme = localStorage.getItem(THEME_NAME);
  let currentTheme = THEMES.includes(storedTheme) ? storedTheme : "light";
  applyTheme(currentTheme);

  function applyTheme(theme) {
    const resolved = THEMES.includes(theme) ? theme : "light";
    document.body.classList.remove("dark", "dark-grey", "solarized-dark", "standard");
    for (const cls of THEME_BODY_CLASSES[resolved]) {
      document.body.classList.add(cls);
    }
    localStorage.setItem(THEME_NAME, resolved);
    currentTheme = resolved;
  }
</script>

<div>
  <h1>
    <a href={WEB_URL}>TopHackerNews</a>
  </h1>
  <div class="btn-group theme-switcher">
    {#each THEMES as theme}
      <button
        type="button"
        class="btn btn-ghost"
        class:btn-primary={currentTheme === theme}
        class:btn-default={currentTheme !== theme}
        on:click={() => applyTheme(theme)}
      >
        {theme}
      </button>
    {/each}
  </div>
  <hr>
</div>
