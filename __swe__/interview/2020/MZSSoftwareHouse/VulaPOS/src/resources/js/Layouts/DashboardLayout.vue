<script setup>
import Sidebar from '@/Components/Sidebar.vue';
import ApplicationLogo from '@/Components/ApplicationLogo.vue';
import Header from '@/Components/Header.vue';
</script>

<template>
    <div class="c-app">
        <div class="c-sidebar c-sidebar-dark c-sidebar-fixed"
            :class="{ 'c-sidebar-lg-show': !$page.url.startsWith('/apps/transactions') }" id="sidebar">
            <div class="c-sidebar-brand d-lg-down-none" style="background: #F05340;">
                <ApplicationLogo class="bg-light rounded shadow-sm p-1" width="35"/>
                <span class="ml-2 font-weight-bold">Vula POS</span>
            </div>
            <Sidebar />
        </div>

        <div class="c-wrapper c-fixed-components">
            <Header />
            <div class="c-body">

                <slot />

                <footer class="c-footer">
                    <div><strong>Vula POS</strong> &copy; 2024 - https://vulapos.com.</div>
                </footer>
            </div>
        </div>
    </div>
</template>
