<script setup>
import { ref } from 'vue';
import { Head } from '@inertiajs/vue3';
//chart
import { BarChart, DoughnutChart } from 'vue-chart-3';
import { Chart, registerables } from "chart.js";
import DashboardLayout from '@/Layouts/DashboardLayout.vue';

//register chart
Chart.register(...registerables);

const props = defineProps({
    count_sales_today: {
        type: Number
    },
    sum_sales_today: {
        type: Number
    },
    sum_profits_today: {
        type: Number
    },
    sales_date: {
        type: Array
    },
    grand_total: {
        type: Array
    },
    product: {
        type: Array
    },
    total: {
        type: Array
    },
    products_limit_stock: {
        type: Array
    },
});

//method random color
function randomBackgroundColor(length) {
    var data = [];
    for (var i = 0; i < length; i++) {
        data.push(getRandomColor());
    }
    return data;
}

console.log(props.products_limit_stock);

//method generate random color
function getRandomColor() {
    var letters = '0123456789ABCDEF'.split('');
    var color = '#';
    for (var i = 0; i < 6; i++) {
        color += letters[Math.floor(Math.random() * 16)];
    }
    return color;
}

//option chart
const options = ref({
    responsive: true,
    plugins: {
        legend: {
            display: false,
        },
        title: {
            display: false,
        },
    },
    beginZero: true
});

//chart sell week
const chartSellWeek = {
    labels: props.sales_date,
    datasets: [{
        data: props.grand_total,
        backgroundColor: randomBackgroundColor(props.sales_date.length),
    },],
};

//chart produk terlaris
const chartBestProduct = {
    labels: props.product,
    datasets: [{
        data: props.total,
        backgroundColor: randomBackgroundColor(5),
    },],
};

</script>

<template>
    <DashboardLayout>

        <Head title="Dashboard" />

        <main class="c-main">
            <div class="container-fluid">
                <div class="fade-in">

                    <div class="row">
                        <div class="col-md-8">
                            <div v-if="hasAnyPermission(['dashboard.sales_chart'])"
                                class="card border-0 rounded-3 shadow border-top-purple">
                                <div class="card-header">
                                    <span class="font-weight-bold"><i class="fa fa-chart-bar"></i> SALES CHART 7
                                        DAYS</span>
                                </div>
                                <div class="card-body">
                                    <BarChart :chartData="chartSellWeek" :options="options" />
                                </div>
                            </div>
                        </div>
                        <div class="col-md-4">

                            <div v-if="hasAnyPermission(['dashboard.sales_today'])"
                                class="card border-0 rounded-3 shadow border-top-info mb-4">
                                <div class="card-header">
                                    <span class="font-weight-bold"><i class="fa fa-chart-line"></i> SALES TODAY</span>
                                </div>
                                <div class="card-body">
                                    <strong>{{ count_sales_today }}</strong> SALES
                                    <hr>
                                    <h5 class="fw-bold">Rp. {{ formatPrice(sum_sales_today) }}</h5>
                                </div>
                            </div>

                            <div v-if="hasAnyPermission(['dashboard.profits_today'])"
                                class="card border-0 rounded-3 shadow border-top-success">
                                <div class="card-header">
                                    <span class="font-weight-bold"><i class="fa fa-chart-bar"></i> PROFITS TODAY</span>
                                </div>
                                <div class="card-body">
                                    <h5 class="fw-bold">Rp. {{ formatPrice(sum_profits_today) }}</h5>
                                </div>
                            </div>

                        </div>
                    </div>

                    <div class="row">
                        <div class="col-md-6">
                            <div v-if="hasAnyPermission(['dashboard.best_selling_product'])"
                                class="card border-0 rounded-3 shadow border-top-warning">
                                <div class="card-header">
                                    <span class="font-weight-bold"><i class="fa fa-chart-pie"></i> BEST SELIING
                                        PRODUCT</span>
                                </div>
                                <div class="card-body">
                                    <DoughnutChart :chartData="chartBestProduct" />
                                </div>
                            </div>
                        </div>
                        <div class="col-md-6">
                            <div v-if="hasAnyPermission(['dashboard.product_stock'])"
                                class="card border-0 rounded-3 shadow border-top-danger">
                                <div class="card-header">
                                    <span class="font-weight-bold"><i class="fa fa-box-open"></i> PRODUCT STOCK</span>
                                </div>
                                <div class="card-body">
                                    <div v-if="products_limit_stock.length > 0">
                                        <ol class="list-group list-group-numbered">
                                            <li v-for="product in products_limit_stock" :key="product.id"
                                                class="list-group-item d-flex justify-content-between align-items-start">
                                                <div class="ms-2 me-auto">
                                                    <div class="fw-bold">{{ product.title }}</div>
                                                    <div class="text-muted">Category : {{ product.category.name }}</div>
                                                </div>
                                                <span class="badge bg-danger rounded-pill">{{ product.stock }}</span>
                                            </li>
                                        </ol>
                                    </div>
                                    <div v-else class="alert alert-danger border-0 shadow rounded-3">
                                        Data Tidak Tersedia!.
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>

                </div>
            </div>
        </main>
    </DashboardLayout>
</template>
