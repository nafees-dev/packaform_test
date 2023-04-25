<template>
  <div class="wrapper">
    <MDBTable class="fs-6 align-middle mb-0 bg-white border">
      <thead>
        <tr>
          <td colspan="6">
            <SearchBar :getAllOrders="getAllOrders" />
          </td>
        </tr>
        <tr>
          <td colspan="6">
            <DatPicker :getAllOrders="getAllOrders" />
          </td>
        </tr>
        <tr>
          <td colspan="6">
            <div class="amount-container">
              <label>Total Amount:</label>
              <label class="fw-bold">{{ totalAmount }}</label>
            </div>
          </td>
        </tr>
        <tr class="bg-light">
          <th>Order name</th>
          <th>Customer Company</th>
          <th>Customer Name</th>
          <th>Order Date</th>
          <th>Delivered Amount</th>
          <th>Total Amount</th>
        </tr>
      </thead>
      <tbody v-if="ordersList.length > 0" class="fs-5">
        <tr v-for="(order, index) in ordersList" :key="index">
          <td>
            <div class="ms-3">
              <p class="fw-bold mb-1">{{ order.order_name }}</p>
              <p class="text-muted mb-0">christmas</p>
            </div>
          </td>
          <td>
            <p class="fw-normal mb-1">{{ order?.Customers?.Customer_Companies?.company_name }}</p>
          </td>
          <td>
            <p class="fw-normal mb-1">{{ order.Customers?.name }}</p>
          </td>
          <td>{{ formatDate(order.created_at) }}</td>
          <td class="fs-4">
            <MDBBadge badge="warning" pill class="d-inline">{{ getTotalDeliveries(order.OrderItems) }}</MDBBadge>
          </td>
          <td class="fs-4">
            <MDBBadge badge="success" pill class="d-inline">{{ getTotalAmount(order.OrderItems) }}</MDBBadge>
          </td>
        </tr>
      </tbody>
    </MDBTable>
    <footer>
      <nav class="fs-4" aria-label="Page navigation example">
        <MDBPagination lg class="justify-content-center">
          <MDBPageNav @click="onChangePage(pageNumber - 1)" prev></MDBPageNav>
          <MDBPageItem @click="onChangePage(n)" v-for="n in totalPages" :key="n">{{ n }}</MDBPageItem>
          <MDBPageNav @click="onChangePage(pageNumber + 1)" next></MDBPageNav>
        </MDBPagination>
      </nav>
    </footer>
  </div>
</template>
  
<script>
import { MDBTable, MDBBadge, MDBPagination, MDBPageNav, MDBPageItem } from 'mdb-vue-ui-kit';
import SearchBar from "../modules/SearchBar.vue"
import DatPicker from "../modules/DatePicker.vue"
import { GET_ORDER } from "../apis/index"

export default {
  name: 'OrderComponent',
  props: {
    msg: String
  },
  components: {
    MDBTable, MDBBadge, SearchBar, DatPicker, MDBPagination, MDBPageNav, MDBPageItem
  },

  data() {
    return {
      ordersList: [],
      totalAmount: 0,
      offset: 5,
      totalOrders: 0,
      pageNumber: 1,
      totalPages: 0,
    }
  },

  created() {
    this.getAllOrders();
  },

  methods: {
    async getAllOrders(createdAt = "", searchStr = "") {
      const header = { "content-type": "application/json" }
      let url = `${GET_ORDER}?createdAt=${createdAt}&search=${searchStr}&limit=${this.offset}&page=${this.pageNumber}`
      const result = await fetch(url, { headers: header });
      if (result.status !== 200) return;
      const response = await result.json();
      this.ordersList = response.data;
      this.totalOrders = response.total
      this.totalPages = response.total / this.offset;
    },

    formatDate(dateString) {
      const date = new Date(dateString);
      const options = { year: 'numeric', month: 'long', day: 'numeric', hour: 'numeric', minute: 'numeric', second: 'numeric', timeZoneName: 'short' };
      return date.toLocaleDateString('en-US', options);
    },

    onChangePage(n) {
      console.log("===", n)
      if (n > 0 && n <= this.totalPages) {
        this.pageNumber = n;
        this.getAllOrders()
      }
    },

    getTotalAmount(orderItem) {
      const orderSum = orderItem?.reduce((acc, item) => {
        const pricePerUnit = parseFloat(item.price_per_unit);
        const quantity = parseInt(item.quantity);

        if (!isNaN(pricePerUnit) && !isNaN(quantity)) {
          const price = pricePerUnit * quantity;
          return acc + price;
        } else {
          return acc;
        }
      }, 0);

      return Number(orderSum).toFixed(2);
    },

    getTotalDeliveries(orderItem) {
      let deliveredAmount = 0;
      if (!orderItem || orderItem.length === 0) return 0
      for (let i = 0; i < orderItem.length; i++) {
        const order = orderItem[i];
        const pricePerUnit = parseFloat(order.price_per_unit) || 0;
        let sum = order?.Deliveries?.reduce((acc, item) => {
          const quantity = parseInt(item.delivered_quantity || 0);
          if (!isNaN(quantity)) {
            const price = pricePerUnit * quantity;
            return acc + price;
          } else {
            return acc;
          }
        }, 0);
        deliveredAmount = (sum || 0) + deliveredAmount;
      }
      return Number(deliveredAmount || 0).toFixed(2)
    }
  }

}
</script>
  
  <!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.wrapper {
  padding-left: 10%;
  padding-right: 10%;
}

.date-wrapper {
  text-align: left;
  align-items: left;
}

.amount-container {
  text-align: left;
}
</style>
  