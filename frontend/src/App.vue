<script setup lang="ts">
import Title from "./components/title.vue";
import Description from "./components/description.vue";
import Form from "./components/form.vue";
</script>

<template>
    <div class="container hero__welcome">
        <Title title="Youtube Downloader" />
        <Description text="Descarga videos de YouTube en formato .mp3 .mp4." />
    </div>
    <div class="container hero__videoForm">
        <Form />
    </div>
</template>

<style scoped>
    .container {
        width: 60%;
        margin: 0 auto;
        color: #fff;
    }
    @media (max-width: 768px) {
        .container {
            width: 90%;
        }
    }
    .hero__welcome {
        text-align: center;
        margin-top: 100px;
    }
    .hero__videoForm  {
        margin-top: 50px;
        text-align: center;
    }
</style>
