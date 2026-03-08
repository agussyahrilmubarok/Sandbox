import "./bootstrap";
import "../css/app.css";
import "sweetalert2/dist/sweetalert2.min.css";

import { createApp, h } from "vue";
import { createInertiaApp } from "@inertiajs/vue3";
import { resolvePageComponent } from "laravel-vite-plugin/inertia-helpers";
import { ZiggyVue } from "../../vendor/tightenco/ziggy";
import VueSweetalert2 from "vue-sweetalert2";

const appName = import.meta.env.VITE_APP_NAME || "VulaPOS";

createInertiaApp({
    title: (title) => `${title} - ${appName}`,
    resolve: (name) =>
        resolvePageComponent(
            `./Pages/${name}.vue`,
            import.meta.glob("./Pages/**/*.vue")
        ),
    setup({ el, App, props, plugin }) {
        return createApp({ render: () => h(App, props) })
            .mixin({
                methods: {
                    hasAnyPermission: function (permissions) {
                        //get permissions from props
                        let allPermissions = this.$page.props.auth.permissions;
                        // console.log('All Permissions from page:', this.$page.props.auth.permissions);
                        // console.log("All Permissions:", allPermissions);
                        // console.log("Checking Permissions:", permissions);

                        let hasPermission = false;
                        // console.log("Has Permission:", hasPermission);
                        permissions.forEach(function (item) {
                            if (allPermissions[item]) hasPermission = true;
                        });

                        return hasPermission;
                    },

                    formatPrice(value) {
                        let val = (value / 1).toFixed(0).replace(".", ",");
                        return val
                            .toString()
                            .replace(/\B(?=(\d{3})+(?!\d))/g, ".");
                    },
                },
            })
            .use(plugin)
            .use(ZiggyVue)
            .use(VueSweetalert2)
            .mount(el);
    },
    progress: {
        color: "#29d",
    },
});
