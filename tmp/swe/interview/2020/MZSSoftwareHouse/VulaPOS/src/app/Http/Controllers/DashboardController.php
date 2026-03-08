<?php

namespace App\Http\Controllers;

use Carbon\Carbon;
use Inertia\Inertia;
use App\Models\Profit;
use App\Models\Product;
use App\Models\Transaction;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\DB;
use App\Http\Controllers\Controller;
use Illuminate\Support\Facades\Log;
use Inertia\Response;

class DashboardController extends Controller
{
    /**
     * Handle the incoming request.
     *
     * @param  \Illuminate\Http\Request  $request
     * @return \Illuminate\Http\Response
     */
    public function index(Request $request): Response
    {
        // Mendapatkan tanggal hari ini
        $day = date('d');

        // Mendapatkan tanggal 7 hari yang lalu
        $week = Carbon::now()->subDays(7);

        // Mendapatkan data penjualan selama 7 hari terakhir
        $chart_sales_week = DB::table('transactions')
            ->addSelect(DB::raw('DATE(created_at) as date, SUM(grand_total) as grand_total'))
            ->where('created_at', '>=', $week)
            ->groupBy('date')
            ->get();

        // Inisialisasi array
        $sales_date = [];
        $grand_total = [];

        // Mengisi array dengan data hasil query
        foreach ($chart_sales_week as $result) {
            $sales_date[] = $result->date;
            $grand_total[] = (int)$result->grand_total;
        }

        // Menghitung penjualan hari ini
        $count_sales_today = Transaction::whereDay('created_at', $day)->count();

        // Menjumlahkan total penjualan hari ini
        $sum_sales_today = Transaction::whereDay('created_at', $day)->sum('grand_total');

        // Menjumlahkan total keuntungan hari ini
        $sum_profits_today = Profit::whereDay('created_at', $day)->sum('total');

        // Mendapatkan produk dengan stok terbatas
        $products_limit_stock = Product::with('category')->where('stock', '<=', 10)->get();

        // Mendapatkan 5 produk terlaris
        $chart_best_products = DB::table('transaction_details')
            ->addSelect(DB::raw('products.title as title, SUM(transaction_details.qty) as total'))
            ->join('products', 'products.id', '=', 'transaction_details.product_id')
            ->groupBy('transaction_details.product_id')
            ->orderBy('total', 'DESC')
            ->limit(5)
            ->get();

        // Inisialisasi array
        $product = [];
        $total = [];

        // Mengisi array dengan data produk terlaris
        foreach ($chart_best_products as $data) {
            $product[] = $data->title;
            $total[] = (int)$data->total;
        }

        Log::info('Received sales_date:', ['sales_date' => $sales_date]);
        // Log::info('Received barcode:', ['product' => $product]);

        // Mengembalikan data untuk tampilan dashboard
        return Inertia::render('Dashboard', [
            'sales_date'           => $sales_date,
            'grand_total'          => $grand_total,
            'count_sales_today'    => (int) $count_sales_today,
            'sum_sales_today'      => (int) $sum_sales_today,
            'sum_profits_today'    => (int) $sum_profits_today,
            'products_limit_stock' => $products_limit_stock,
            'product'              => $product,
            'total'                => $total
        ]);
    }
}
