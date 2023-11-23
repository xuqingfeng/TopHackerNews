<script>
  const WEB_URL = import.meta.env.VITE_WEB_URL;
  const THEME_NAME = "theme";
  const storedTheme = localStorage.getItem(THEME_NAME);
  changeTheme(storedTheme);

  // TODO: find a native 'svelte' way to handle this
  function changeTheme(theme) {
    switch (theme) {
      case "light":
        document.body.classList.remove("dark");
        document.body.classList.remove("dark-grey");
        document.body.classList.remove("solarized-dark");
        document.body.classList.remove("standard");

        localStorage.setItem(THEME_NAME, theme);
        break;
      case "dark":
        document.body.classList.add("dark");
        document.body.classList.remove("dark-grey");
        document.body.classList.remove("solarized-dark");
        document.body.classList.remove("standard");

        localStorage.setItem(THEME_NAME, theme);
        break;
      case "dark-grey":
        document.body.classList.add("dark-grey");
        document.body.classList.remove("dark");
        document.body.classList.remove("solarized-dark");
        document.body.classList.remove("standard");

        localStorage.setItem(THEME_NAME, theme);
        break;
      case "solarized-dark":
        document.body.classList.add("solarized-dark");
        document.body.classList.remove("dark");
        document.body.classList.remove("dark-grey");
        document.body.classList.remove("standard");

        localStorage.setItem(THEME_NAME, theme);
        break;
      default:
        document.body.classList.remove("dark");
        document.body.classList.remove("dark-grey");
        document.body.classList.remove("solarized-dark");
        document.body.classList.remove("standard");

        localStorage.setItem(THEME_NAME, "light");
    }
  }
</script>

<div>
  <h1>
    <a href={WEB_URL}>TopHackerNews</a>
  </h1>
  <a on:click={() => changeTheme("light")}>light</a>
  <a on:click={() => changeTheme("dark")}>dark</a>
  <a on:click={() => changeTheme("dark-grey")}>dark-grey</a>
  <a on:click={() => changeTheme("solarized-dark")}>solarized-dark</a>
  <hr>
</div>
