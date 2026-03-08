<?php

namespace App\Http\Controllers;

use App\Models\Cart;
use App\Models\Customer;
use App\Models\Product;
use App\Models\Transaction;
use Illuminate\Contracts\View\View;
use Illuminate\Http\RedirectResponse;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Log;
use Illuminate\Support\Str;
use Inertia\Inertia;
use Inertia\Response;

class TransactionController extends Controller
{
    /**
     * Display a listing of the resource.
     */
    public function index(): Response
    {
        $carts = Cart::with('product')->where('cashier_id', auth()->user()->id)
            ->latest()->get();

        $customers = Customer::latest()->get();

        $products = Product::all();

        return Inertia::render('Transactions/Index', [
            'carts'         => $carts,
            'carts_total'   => $carts->sum('price'),
            'customers'     => $customers,
            'products'      => $products,
        ]);
    }

    /**
     * searchProduct
     *
     * @param  mixed $request
     * @return void
     */
    public function searchProduct(Request $request)
    {
        $product = Product::where('barcode', $request->barcode)->first();

        // Log::info('Received barcode:', ['barcode' => $request->barcode]);
        // Log::info('Received barcode:', ['product' => $product]);

        if ($product) {
            return response()->json([
                'success' => true,
                'data'    => $product
            ]);
        }
    }

    /**
     * addToCart
     *
     * @param  mixed $request
     * @return void
     */
    public function addToCart(Request $request): RedirectResponse
    {
        $product = Product::find($request->product_id);
        if (!$product) {
            return redirect()->back()
                ->with('error', 'Product not found.');
        }
        if ($product->stock < $request->qty) {
            return redirect()->back()
                ->with('error', 'Out of Stock Product!');
        }

        $cart = Cart::where('product_id', $request->product_id)
            ->where('cashier_id', auth()->user()->id)
            ->first();

        if ($cart) {
            $cart->increment('qty', $request->qty);
            $cart->price = $cart->product->sell_price * $cart->qty;
            $cart->save();
        } else {
            Cart::create([
                'cashier_id'    => auth()->user()->id,
                'product_id'    => $request->product_id,
                'qty'           => $request->qty,
                'price'         => $product->sell_price * $request->qty,
            ]);
        }

        return redirect()->route('dashboard.transactions.index')
            ->with('success', 'Product added to cart successfully!');
    }

    /**
     * destroyCart
     *
     * @param  mixed $request
     * @return void
     */
    public function destroyCart(Request $request): RedirectResponse
    {
        $cart = Cart::find($request->cart_id);

        if (!$cart) {
            return redirect()->back()->with('error', 'Cart item not found.');
        }

        $cart->delete();

        return redirect()->back()->with('success', 'Product removed successfully!');
    }

    /**
     * Store a newly created resource in storage.
     */
    public function store(Request $request)
    {
        /**
         * algorithm generate no invoice
         */
        $length = 10;
        $random = '';
        for ($i = 0; $i < $length; $i++) {
            $random .= rand(0, 1) ? rand(0, 9) : chr(rand(ord('a'), ord('z')));
        }

        //generate no invoice
        $invoice = 'TRX-' . Str::upper($random);

        //insert transaction
        $transaction = Transaction::create([
            'cashier_id'    => auth()->user()->id,
            'customer_id'   => $request->customer_id,
            'invoice'       => $invoice,
            'cash'          => $request->cash,
            'change'        => $request->change,
            'discount'      => $request->discount,
            'grand_total'   => $request->grand_total,
        ]);

        //get carts
        $carts = Cart::where('cashier_id', auth()->user()->id)->get();

        //insert transaction detail
        foreach ($carts as $cart) {

            //insert transaction detail
            $transaction->details()->create([
                'transaction_id'    => $transaction->id,
                'product_id'        => $cart->product_id,
                'qty'               => $cart->qty,
                'price'             => $cart->price,
            ]);

            //get price
            $total_buy_price  = $cart->product->buy_price * $cart->qty;
            $total_sell_price = $cart->product->sell_price * $cart->qty;

            //get profits
            $profits = $total_sell_price - $total_buy_price;

            //insert profit
            $transaction->profits()->create([
                'transaction_id'    => $transaction->id,
                'total'             => $profits,
            ]);

            //update stock product
            $product = Product::find($cart->product_id);
            $product->stock = $product->stock - $cart->qty;
            $product->save();
        }

        //delete carts
        Cart::where('cashier_id', auth()->user()->id)->delete();

        return response()->json([
            'success' => true,
            'data'    => $transaction
        ]);
    }

    /**
     * print
     *
     * @param  mixed $request
     * @return void
     */
    public function print(Request $request): View
    {
        $transaction = Transaction::with('details.product', 'cashier', 'customer')
            ->where('invoice', $request->invoice)->firstOrFail();

        return view('print.nota', compact('transaction'));
    }
}
