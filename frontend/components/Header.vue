<template>
  <div class="header">
    <div class="logo">
      <NuxtLink :to="{ name: 'index' }">duck</NuxtLink>
    </div>

    <div class="user">
      <div class="info">
        <div class="list">
          <NuxtLink :to="{ name: 'login' }">Log in</NuxtLink>
          <NuxtLink :to="{ name: 'register' }">Register</NuxtLink>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
</script> 