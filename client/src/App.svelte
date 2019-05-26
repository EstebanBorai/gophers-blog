<script>
	export let name;

  let users = [{"id":"13067ca2-40bf-443e-b3fc-7281c011964f","firstName":"John","lastName":"Doe","userName":"johndoe"}];

  fetch('http://localhost:8080/api/users').then((res) => res.json()).then((json) => {
    users = JSON.stringify(json);
  }).catch((err) => {
    // cors issue here
    console.error(err);
  });
</script>

<style>
  #app-container {
    background: black;
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
  }

	#app-main {
    background: #ffffff;
    width: 414px;
    height: 680px;
  }

  #app-main header {
    height: 65px;
    background: #0f55e0;
    color: #ffffff;
  }

  #app-main header h1 {
    margin: 0;
    padding: 0;
  }
</style>

<div id="app-container">
  <main id="app-main">
    <header>
      <h1>Songs Share</h1>
    </header>
    <h4>All of our users!</h4>
    {#each users as user}
      <h2>{user.firstName}</h2>
    {/each}
  </main>
</div>
