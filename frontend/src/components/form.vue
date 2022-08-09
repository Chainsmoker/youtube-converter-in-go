<template>
        <form method="post" @submit.prevent="verificarForm" id="form">
            <div class="form__input">
                <input type="text" id="id_videoURL" v-model="videoURL" placeholder="Ingresa la URL del video" required />
            </div>
            <div class="form__button">
                <button type="submit">Convertir</button>
            </div>
        </form>
</template>

<script lang="ts">
export default {
    methods: {
        verificarForm: async function () {
            const input = document.getElementById("id_videoURL") as HTMLInputElement;
            if ((this as any).videoURL.length < 6 || (this as any).videoURL.includes("youtube.com") == false) {
                input.style.borderColor = "red";
                alert('La URL no es valida, incluye el dominio "youtube.com"');
            } else {
                input.style.borderColor = "green";
                // En entorno local la URL debe ser localhost:port/download
                const response = await fetch('/download', {
                    method: 'POST',
                    headers: {
                        'Accept': 'application/json',
                        'Content-Type': 'application/json',
                    },
                    body: JSON.stringify({
                        url: (this as any).videoURL,
                    })
                });
                const data = await response.json();
                if (data){
                    location.href = data.download_url;
                }

            }
        }
    },
}
</script>

<style scoped>
    input {
        width: 90%;
        height: 50px;
        border-radius: 10px;
        margin-bottom: 20px;
        padding-left: 10px;
        border: 2px solid;
    }
    input:focus {
        outline: none;
    }
    button {
        padding: 10px 20px;
        border-radius: 5px;
        outline: none;
        border: none;
        background-color: rgb(255, 127, 80);
        color: #fff;
    }
    button:hover {
        background-color: rgb(255, 69, 0);
    }
</style>
