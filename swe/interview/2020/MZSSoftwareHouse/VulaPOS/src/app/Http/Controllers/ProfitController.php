<?php

namespace App\Http\Controllers;

use App\Exports\ProfitsExport;
use App\Models\Profit;
use Barryvdh\DomPDF\Facade\Pdf;
use Illuminate\Http\Request;
use Illuminate\Http\Response as HttpResponse;
use Inertia\Inertia;
use Inertia\Response;
use Maatwebsite\Excel\Facades\Excel;
use Symfony\Component\HttpFoundation\BinaryFileResponse;

class ProfitController extends Controller
{
    /**
     * index
     *
     * @return void
     */
    public function index(): Response
    {
        return Inertia::render('Profits/Index');
    }

    /**
     * filter
     *
     * @param  mixed $request
     * @return void
     */
    public function filter(Request $request): Response
    {
        $request->validate([
            'start_date' => ['required'],
            'end_date' => ['required']
        ]);

        //get data profits by range date
        $profits = Profit::with('transaction')->whereDate('created_at', '>=', $request->start_date)->whereDate('created_at', '<=', $request->end_date)->get();

        //get total profit by range date
        $total = Profit::whereDate('created_at', '>=', $request->start_date)->whereDate('created_at', '<=', $request->end_date)->sum('total');

        return Inertia::render('Profits/Index', [
            'profits'   => $profits,
            'total'     => (int) $total
        ]);
    }

    /**
     * export
     *
     * @param  mixed $request
     * @return void
     */
    public function export(Request $request): BinaryFileResponse
    {
        return Excel::download(new ProfitsExport($request->start_date, $request->end_date), 'profits : ' . $request->start_date . ' — ' . $request->end_date . '.xlsx');
    }

    /**
     * pdf
     *
     * @param  mixed $request
     * @return void
     */
    public function pdf(Request $request): HttpResponse
    {
        //get data profits by range date
        $profits = Profit::with('transaction')->whereDate('created_at', '>=', $request->start_date)->whereDate('created_at', '<=', $request->end_date)->get();

        //get total profit by range date
        $total = Profit::whereDate('created_at', '>=', $request->start_date)->whereDate('created_at', '<=', $request->end_date)->sum('total');

        //load view PDF with data
        $pdf = PDF::loadView('exports.profits', compact('profits', 'total'));

        //download PDF
        return $pdf->download('profits : ' . $request->start_date . ' — ' . $request->end_date . '.pdf');
    }
}
