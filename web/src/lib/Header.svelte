<script>
  const WEB_URL = import.meta.env.VITE_WEB_URL;

  // TODO: find a native 'svelte' way to handle this
  function changeTheme(theme) {
    switch (theme) {
      case "light":
        document.getElementById("body").classList.remove("dark");
        document.getElementById("body").classList.remove("dark-grey");
        document.getElementById("body").classList.remove("solarized-dark");
        document.getElementById("body").classList.remove("standard");
        break;
      case "dark":
        document.getElementById("body").classList.add("dark");
        document.getElementById("body").classList.remove("dark-grey");
        document.getElementById("body").classList.remove("solarized-dark");
        document.getElementById("body").classList.remove("standard");
        break;
      case "dark-grey":
        document.getElementById("body").classList.add("dark-grey");
        document.getElementById("body").classList.remove("dark");
        document.getElementById("body").classList.remove("solarized-dark");
        document.getElementById("body").classList.remove("standard");
        break;
      case "solarized-dark":
        document.getElementById("body").classList.add("solarized-dark");
        document.getElementById("body").classList.remove("dark");
        document.getElementById("body").classList.remove("dark-grey");
        document.getElementById("body").classList.remove("standard");
        break;
      default:
        document.getElementById("body").classList.remove("dark");
        document.getElementById("body").classList.remove("dark-grey");
        document.getElementById("body").classList.remove("solarized-dark");
        document.getElementById("body").classList.remove("standard");
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
