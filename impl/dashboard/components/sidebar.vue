<template>
  <nav class="flex flex-col justify-between">
    <ul>
      <div>
        <li>
          <NuxtLink to="/account">account</NuxtLink>
        </li>
        <li><NuxtLink to="/devices">devices</NuxtLink></li>
        <li><NuxtLink to="/groups">groups</NuxtLink></li>
        <li>
          <NuxtLink to="/deployments">deployments</NuxtLink>
        </li>
        <li>
          <NuxtLink to="/recipes">Recipes</NuxtLink>
        </li>
      </div>
      <li>
        <NuxtLink to="logout">logout</NuxtLink>
      </li>
    </ul>
  </nav>
</template>

<style></style>
