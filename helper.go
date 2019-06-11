package ovo

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"time"
)

type Stop struct {
	error
}

func GetDetailResponse(lang, responseCode string) (string, string) {
	if(lang == "id") {
		switch responseCode {
			case "00":
				return "Berhasil / Diterima", "Transaksi berhasil diterima"
            case "8":
                return "Kode Voucher Sudah Terpakai", "Gagal melakukan transaksi, voucher sudah digunakan"
            case "11":
                return "Jenis Akun Tidak Valid", "Tipe akun tidak valid"
            case "13":
                return "Jumlah Tidak Valid", "Menggunakan jumlah kurang dari atau sama dengan nol"
            case "14":
                return "Nomor Ponsel / ID OVO Tidak Valid", "Nomor telepon / ID OVO tidak ditemukan di Sistem OVO"
            case "16":
                return "Kartu Tidak Valid", "Nomor kartu tidak valid"
            case "17":
                return "Transaksi Ditolak", "Pengguna OVO membatalkan pembayaran menggunakan Aplikasi OVO"
            case "18":
                return "Kode Voucher Tidak Valid", "Gagal melakukan transaksi, kode voucher tidak valid"
			case "25":
				return "Transaksi Tidak Ditemukan", "Status pembayaran tidak ditemukan pada saat mengecek status pembayaran"
            case "26":
                return "Transaksi Gagal", "Gagal mengkonfirmasi pembayaran ke Aplikasi OVO"
            case "30":
                return "Kesalahan Format", "Spesifikasi ISO tidak valid"
            case "40":
                return "Transaksi Gagal", "Kesalahan pihak ke-3, termasuk gagal mengurangi / menambah uang elektronik"
            case "44":
                return "Kode Refund Kadaluwarsa", "Kode refund sudah tidak berlaku (selama batas waktu)"
            case "51":
                return "Dana Tidak Memadai", "Dana dalam kartu tidak cukup untuk melakukan pembayaran"
            case "54":
                return "Transaksi Kadaluwarsa. (Melebihi 7 hari)", "Detail transaksi sudah kadaluwarsa pada saat mengecek status pembayaran"
            case "56":
                return "Kartu Diblokir. Silakan Hubungi 1500696", "Kartu diblokir, tidak dapat memproses transaksi kartu"
            case "57":
                return "Transaksi Gagal", "Tidak memenuhi syarat untuk melakukan refund"
            case "58":
                return "Transaksi Tidak Diizinkan", "Tidak berlaku transaksi di merchant"
            case "61":
                return "Transaksi Melebihi Batas", "Jumlah / hitungan melebihi batas yang ditetapkan oleh pengguna"
            case "63":
                return "Pelanggaran Keamanan", "Otentikasi gagal"
            case "64":
                return "Akun Diblokir. Silahkan Hubungi 1500696", "Akun diblokir, tidak dapat memproses transaksi"
            case "65":
                return "Transaksi Gagal", "Transaksi melebihi batas"
            case "66":
                return "Kode Tidak Valid", "Kode penarikan dana salah"
            case "67":
                return "Transaksi Dibawah Batas", "Jumlah transaksi kurang dari pembayaran minimum"
			case "68":
				return "Transaksi Tertunda / Kadaluwarsa", "OVO Wallet telat memberikan respon ke OVO JPOS"
			case "73":
				return "Transaksi Telah Dikembalikan", "Transaksi sudah dikembalikan"
			case "96":
				return "Kode Pemrosesan Tidak Valid","Kode pemrosesan tidak valid pada saat mengecek status pembayaran"
            case "ER":
                return "Kegagalan Sistem", "Ada kesalahan dalam Sistem OVO"
            case "EB":
                return "Terminal Diblokir", "Tidak terdaftar / diblokir terminal"
            default:
                return "Tidak Ditemukan", "Pembayaran kadaluwarsa / pelanggan mengabaikan transaksi"
		}
	} else {
		switch responseCode {
			case "00":
				return "Success / Approved", "Transaction successfully approved"
            case "8":
                return "Voucher Code Already Used", "Failed to redeem deal, voucher already used"
            case "11":
                return "Invalid Account Type", "Invalid user level, unable to proceed cash withdrawal"
            case "13":
                return "Invalid Amount", "Using amount less than equal zero"
            case "14":
                return "Invalid Mobile Number/OVO ID", "Phone number / OVO ID not found in OVO System"
            case "16":
                return "Invalid Card", "Invalid card number"
            case "17":
                return "Transaction Declined", "OVO user cancelled payment using OVO Apps"
            case "18":
                return "Invalid Voucher Code", "Failed to redeem deal, invalid voucher code"
			case "25":
				return "Transaction Not Found", "Payment status not found when called by Check Payment Status API"
            case "26":
                return "Transaction Failed", "Failed to push payment confirmation to OVO Apps"
            case "30":
                return "Format Error", "Invalid ISO Specs"
            case "40":
                return "Transaction Failed", "Error in 3rd party, including failed to deduct / topup e-money"
            case "44":
                return "Expired Refund Code", "Refund code is no longer valid (over the time limit)"
            case "51":
                return "Insufficient Fund", "Fund in card is not enough to make payment"
			case "54":
				return "Transaction Expired (More than 7 days)", "Transaction details already expired when API check payment status called"
            case "56":
                return "Card Blocked. Please Call 1500696", "Card is blocked, unable to process card transaction"
            case "57":
                return "Transaction Failed", "Not eligible to do refund, siloam balance = 0"
            case "58":
                return "Transaction Not Allowed", "Invalid transaction in merchant/terminal"
            case "61":
                return "Exceed transaction limit", "Amount/count exceed limit, set by user"
            case "63":
                return "Security Violation", "Authentication failed"
            case "64":
                return "Account Blocked. Please Call 1500696", "Account is blocked, unable to process transaction"
            case "65":
                return "Transaction Failed", "Limit transaction exceeded, limit on count or amount"
            case "66":
                return "Invalid Code", "Wrong withdrawal code"
            case "67":
                return "Below Transaction Limit", "The transaction amount is less than the minimum payment"
			case "68":
				return "Transaction Pending / Timeout", "OVO Wallet late to give respond to OVO JPOS"
			case "73":
				return "Transaction Has Been Reversed", "Transaction has been reversed by API Reversal Push to Pay in Check Payment Status API"
			case "96":
				return "Invalid Processing Code", "Invalid processing code inputted during Call API Check Payment Status"
            case "ER":
                return "System Failure", "There is an error in OVO System"
            case "EB":
                return "Terminal Blocked", "Not registered / blocked terminal"
            default:
                return "Not Found", "Expired payment push / customer ignore the transaction"
		}
	}
}

func EncodeHMACSHA256(appId, random, key string) string {
	payload := appId + random
    h := hmac.New(sha256.New, []byte(key))
    h.Write([]byte(payload))

    bs := h.Sum(nil)
    result := hex.EncodeToString(bs)
    return result
}

func GetType(name string) string {
    switch name {
	    case "create", "void":
	        return "0200"
	    case "reversal":
	        return "0400"
	    case "status":
	        return "0100"
	    default:
	        return ""
    }
}

func GetProcessingCode(name string) string {
    switch name {
	    case "create", "reversal", "status":
	        return "040000"
	    case "void":
	        return "020040"
	    default:
	        return ""
    }
}

// Copy this from https://upgear.io/blog/simple-golang-retry-function/
func Retry(attempts int, sleep time.Duration, f func() error) error {
	if err := f(); err != nil {
		if s, ok := err.(Stop); ok {
			// Return the original error for later checking
			return s.error
		}

		if attempts--; attempts > 0 {
			time.Sleep(sleep)
			return Retry(attempts, sleep, f)
		}
		return err
	}

	return nil
}