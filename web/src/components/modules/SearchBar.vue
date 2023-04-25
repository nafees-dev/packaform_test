<template>
    <div class="header-container">
        <MDBIcon icon="search" />
        <span class="fs-5">Search</span>
        <div class="searchOrder">
            <MDBInput v-model="searchTerm" :formOutline="false" placeholder="Type here ..." aria-label="Search"
                aria-describedby="button" />
        </div>
    </div>
</template>
    
<script>
import { MDBInput, MDBIcon } from 'mdb-vue-ui-kit';
import { debounce } from 'lodash';

export default {
    name: 'SearchBar',
    props: {
        getAllOrders: Function
    },
    components: {
        MDBInput, MDBIcon
    },
    data() {
        return {
            searchTerm: '',
            items: [ /* your list of items to search */],
        };
    },
    computed: {
        filteredItems() {
            return this.items.filter((item) => {
                // filter items based on search term
                return item.name.toLowerCase().includes(this.searchTerm.toLowerCase());
            });
        },
    },
    watch: {
        // use debounce to delay search until user stops typing
        searchTerm: debounce(function (newSearchTerm) {
            console.log('Searching for:', newSearchTerm);
            this.getAllOrders("", newSearchTerm)
        }, 500),
    },
}
</script>
    

<style scoped>
.searchOrder {
    width: 100%;
    margin-left: 20px;
}

.header-container {
    display: flex;
    flex-direction: row;
    width: 100%;
}
</style>
    