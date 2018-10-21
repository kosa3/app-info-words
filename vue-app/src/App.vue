<template>
  <div class="card-columns">
      <div class="card text-center" v-for="item in items">
          <div class="card-body" :data-word-id="item.id">
              <h4 class="card-title">{{ item.name }}</h4>
              <p class="card-text">{{ item.description }}</p>
              <p class="card-text"><small class="text-muted">{{ item.code }}</small></p>
          </div>
      </div>
  </div>
</template>

<script>

"use strict";

export default {
    data() {
        return {
            items: this.fetchWords()
        }
    },
    methods: {
        fetchWords() {
            fetch('http://localhost:8060/api/words').then(response =>  {
                return response.json();
            }).then(res => {
                this.items = res;
                return res;
            });
        }
    }
}
</script>
