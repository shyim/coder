import Button from "@material-ui/core/Button"
import { makeStyles, useTheme } from "@material-ui/core/styles"
import useMediaQuery from "@material-ui/core/useMediaQuery"
import KeyboardArrowLeft from "@material-ui/icons/KeyboardArrowLeft"
import KeyboardArrowRight from "@material-ui/icons/KeyboardArrowRight"
import { ChooseOne, Cond } from "components/Conditionals/ChooseOne"
import { Maybe } from "components/Conditionals/Maybe"
import { CSSProperties } from "react"
import { useSearchParams } from "react-router-dom"
import { PageButton } from "./PageButton"
import { buildPagedList, DEFAULT_RECORDS_PER_PAGE } from "./utils"

export type PaginationWidgetProps = {
  prevLabel: string
  nextLabel: string
  onPageChange: (offset: number, limit: number) => void
  numRecordsPerPage?: number
  numRecords?: number
  containerStyle?: CSSProperties
}

export const PaginationWidget = ({
  prevLabel,
  nextLabel,
  onPageChange,
  numRecords,
  numRecordsPerPage = DEFAULT_RECORDS_PER_PAGE,
  containerStyle,
}: PaginationWidgetProps): JSX.Element | null => {
  const theme = useTheme()
  const isMobile = useMediaQuery(theme.breakpoints.down("sm"))
  const styles = useStyles()

  const [searchParams, setSearchParams] = useSearchParams()
  const currentPage = searchParams.get("page")
    ? Number(searchParams.get("page"))
    : 1

  const numPages = numRecords ? Math.ceil(numRecords / numRecordsPerPage) : 0
  const firstPageActive = currentPage === 1 && numPages !== 0
  const lastPageActive = currentPage === numPages && numPages !== 0

  const changePage = (newPage: number) => {
    // change the page in the url
    setSearchParams({ page: newPage.toString() })
    // fetch new page of records
    const offset = (newPage - 1) * numRecordsPerPage
    onPageChange(offset, numRecordsPerPage)
  }

  // No need to display any pagination if we know the number of pages is 1
  if (numPages === 1 || numRecords === 0) {
    return null
  }

  return (
    <div style={containerStyle} className={styles.defaultContainerStyles}>
      <Button
        className={styles.prevLabelStyles}
        aria-label="Previous page"
        disabled={firstPageActive}
        onClick={() => changePage(currentPage - 1)}
      >
        <KeyboardArrowLeft />
        <div>{prevLabel}</div>
      </Button>
      <Maybe condition={numPages > 0}>
        <ChooseOne>
          <Cond condition={isMobile}>
            <PageButton
              activePage={currentPage}
              page={currentPage}
              numPages={numPages}
            />
          </Cond>
          <Cond>
            {buildPagedList(numPages, currentPage).map((page) =>
              typeof page !== "number" ? (
                <PageButton key={`Page${page}`} placeholder="..." disabled />
              ) : (
                <PageButton
                  key={`Page${page}`}
                  activePage={currentPage}
                  page={page}
                  numPages={numPages}
                  onPageClick={() => changePage(page)}
                />
              ),
            )}
          </Cond>
        </ChooseOne>
      </Maybe>
      <Button
        aria-label="Next page"
        disabled={lastPageActive}
        onClick={() => changePage(currentPage + 1)}
      >
        <div>{nextLabel}</div>
        <KeyboardArrowRight />
      </Button>
    </div>
  )
}

const useStyles = makeStyles((theme) => ({
  defaultContainerStyles: {
    justifyContent: "center",
    alignItems: "center",
    display: "flex",
    flexDirection: "row",
    padding: "20px",
  },

  prevLabelStyles: {
    marginRight: `${theme.spacing(0.5)}px`,
  },
}))
