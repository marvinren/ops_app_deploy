<template>
  <div ref="code_container" style="height:500px"></div>
</template>
<script lang="ts">
import * as monaco from 'monaco-editor'

export default {
  props: ["confName", "confScript"],
  data: ()=>({
    editor:null,
  }),
  mounted(){

    let option:any = {
      'value': this.confScript,
      'language': 'yaml',
      'automaticLayout': true,
      'minimap': {
        'enabled': false
      }
    };
    this.editor = monaco.editor.create(this.$refs.code_container, option)
  }
}
</script>