from __future__ import print_function

import sys
import os
import dropbox
import glob
from dropbox.files import WriteMode
from dropbox.exceptions import ApiError, AuthError


SCRIPT_FOLDER = os.path.dirname(os.path.abspath(__file__))
CSV_FOLDER = os.path.dirname(SCRIPT_FOLDER) + "/csvs/"

# (https://blogs.dropbox.com/developers/2014/05/generate-an-access-token-for-your-own-account/)
dbx = dropbox.Dropbox(os.environ["DROPBOX_ACCESS_TOKEN"])

# check token
try:
    dbx.users_get_current_account()
except AuthError:
    sys.exit("Invalid access token")


def upload_file(filename):
    with open(CSV_FOLDER + filename, "rb") as f:
        try:
            print("Uploading file: {}".format(filename))
            dbx.files_upload(f.read(), "/{}".format(filename), mode=WriteMode("overwrite"))
        except ApiError as err:
            # This checks for the specific error where a user doesn't have
            # enough Dropbox space quota to upload this file
            if (err.error.is_path() and
                    err.error.get_path().error.is_insufficient_space()):
                sys.exit("ERROR: Cannot back up; insufficient space.")
            elif err.user_message_text:
                print(err.user_message_text)
                sys.exit()
            else:
                print(err)
                sys.exit()


def list_dropbox_exists_filenames(filenames, metadata=None):
    if metadata is not None:
        if metadata.has_more:
            list_folder_result = dbx.files_list_folder_continue(metadata.cursor)
        else:
            return filenames
    else:
        list_folder_result = dbx.files_list_folder("")

    filenames += [entry.name for entry in list_folder_result.entries]
    return list_dropbox_exists_filenames(filenames, list_folder_result)

if __name__ == "__main__":
    exists_files = set(list_dropbox_exists_filenames([]))
    local_files = set(os.path.basename(f) for f in glob.glob(CSV_FOLDER + "*.csv"))

    waiting_to_upload_files = (local_files - exists_files)
    print("Uploading {} files.".format(len(waiting_to_upload_files)))
    for f in waiting_to_upload_files:
        upload_file(f)