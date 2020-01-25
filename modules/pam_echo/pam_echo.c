#include <stdio.h>
#include <stdlib.h>

#ifdef __APPLE__
  #include <security/pam_appl.h>
  #include <security/pam_modules.h>
#else
  #include <security/pam_appl.h>
  #include <security/pam_modules.h>
  #include <security/pam_ext.h>
#endif

int pam_sm_acct_mgmt(pam_handle_t *pamh, int flags, int argc, const char **argv)
{
    int pam_err;
	const char **login_context;
    pam_err = pam_get_item(pamh, PAM_RUSER, (const void **)login_context);
    if (pam_err < 0) {
        const char *err_msg = pam_strerror(pamh, pam_err);
        pam_info(pamh, "pam_sm_acct_mgmt failed: %s", err_msg);
        return PAM_SYSTEM_ERR;
    }

    pam_info(pamh, "%s", *login_context);
    return PAM_SUCCESS;
}

int pam_sm_open_session(pam_handle_t *pamh, int flags, int argc, const char **argv)
{
    int pam_err;
	const char **login_context;
    pam_err = pam_get_item(pamh, PAM_RUSER, (const void **)login_context);
    if (pam_err < 0) {
        const char *err_msg = pam_strerror(pamh, pam_err);
        pam_info(pamh, "pam_sm_open_session failed: %s", err_msg);
        return PAM_SYSTEM_ERR;
    }

    pam_info(pamh, "%s", *login_context);
    return PAM_SUCCESS;
}

int pam_sm_close_session (pam_handle_t *pamh, int flags, int argc, const char **argv)
{
    return PAM_SUCCESS;
}
