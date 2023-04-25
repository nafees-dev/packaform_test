<template>
    <div class="date-wrapper">
        <label class="fw-bold">Created Date</label>
        <VueDatepickerUi lang="en" range show-clear-button v-model="selectedDate" @input="onReset" />
    </div>
</template>
    
<script>
import 'vue-datepicker-ui/lib/vuedatepickerui.css';
import VueDatepickerUi from 'vue-datepicker-ui';

export default {
    name: 'DatePicker',
    props: {
        getAllOrders: Function,
    },
    components: {
        VueDatepickerUi
    },

    data() {
        return {
            selectedDate: [],
        }
    },
    watch: {
        selectedDate(newSelectedDate) {
            if (newSelectedDate[0] && newSelectedDate[1]) {
                this.getAllOrders(JSON.stringify(newSelectedDate))
            }
        }
    },

    methods: {
        onReset() {
            console.log("Rest")
        }
    }
}
</script>
    
    <!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.date-wrapper {
    display: flex;
    flex-direction: column;
    width: 100%;
}
</style>
    