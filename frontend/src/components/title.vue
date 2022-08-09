<template>
  <h1>{{ title }}</h1>
</template>

<script lang="ts">
export default {
    props: {
        title: {
            type: String,
        }
    }
}
</script>

<style scoped>
    h1 {
        font-size: 4rem;
    }

    @media (max-width: 768px) {
        h1 {
            font-size: 2rem;
        }
    }
</style>