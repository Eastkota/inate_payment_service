package helpers

func MapAppleToBankCode(appleStatus int) string {
    switch appleStatus {
    case 0:
        return "Approved" // Approved
    case 21002:
        return "Format Error" // Format Error
    case 21007, 21008:
        return "Invalid Transaction (Sandbox/Prod mismatch)" // Invalid Transaction (Sandbox/Prod mismatch)
    default:
        return "Internal Error" // Internal Error
    }
}